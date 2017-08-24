package com.youstars.sgsocket;

import java.io.UnsupportedEncodingException;
import java.net.URI;
import java.net.URLDecoder;
import java.util.Iterator;
import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.youstars.sgsocket.vo.CmdVo;
import com.youstars.sgsocket.vo.InvalidTokenResponseVo;
import com.youstars.sgsocket.vo.NotificationVO;
import com.youstars.sgsocket.vo.ResponseVo;
import com.youstars.sgsocket.vo.SuccessResponseVo;

import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpRequest;
import io.netty.handler.codec.http.HttpHeaderNames;
import io.netty.handler.codec.http.HttpHeaders;
import io.netty.handler.codec.http.HttpMethod;
import io.netty.handler.codec.http.HttpResponse;
import io.netty.handler.codec.http.HttpResponseStatus;
import io.netty.handler.codec.http.websocketx.TextWebSocketFrame;
import io.netty.handler.codec.http.websocketx.WebSocketServerHandshaker;
import io.netty.handler.codec.http.websocketx.WebSocketServerHandshakerFactory;

public class HttpRequestHandler extends SimpleChannelInboundHandler<FullHttpRequest> {
    
    private static Logger logger = LoggerFactory.getLogger(HttpRequestHandler.class);
    
    private static final String WEBSOCKET_UPGRADE = "websocket";
    private static final String WEBSOCKET_CONNECTION = "Upgrade";
    //private static final String WEBSOCKET_URI_ROOT_PATTERN = "ws://%s:%d";

    private WebSocketServerHandshaker handshaker;

    @Override
    protected void channelRead0(ChannelHandlerContext ctx, FullHttpRequest request) throws Exception {

        //logger.debug("request uri: " + request.uri());
        
        logger.debug("\n" + request.toString());
        
        Iterator<Entry<String, String>> iter = request.headers().iteratorAsString();

        String accessToken = null;
        while (iter.hasNext()) {
            Entry<String, String> entry = iter.next();
            //logger.debug("header: key={},value={}", entry.getKey(), entry.getValue());
            if (entry.getKey().equals("Sec-WebSocket-Protocol")) {
                accessToken = entry.getValue();
            }
            if (entry.getKey().equals("Access-Token")) {
                accessToken = entry.getValue();
                break;
            }
        }
        
        if (request.uri().equals("/ws")) {

            if (isWebSocketUpgrade(request)) {

                // 构造握手响应返回
                WebSocketServerHandshakerFactory ws = new WebSocketServerHandshakerFactory("",
                        request.headers().getAsString("Sec-WebSocket-Protocol"), false);
                handshaker = ws.newHandshaker(request);
                if (handshaker == null) {
                    // 版本不支持
                    WebSocketServerHandshakerFactory.sendUnsupportedVersionResponse(ctx.channel());
                    ctx.channel().close();
                    return;
                }
                handshaker.handshake(ctx.channel(), request);

                if (accessToken == null || accessToken.isEmpty()) {
                    logger.debug("not found access token ....");
                    ctx.channel().writeAndFlush(new TextWebSocketFrame(new CmdVo(Command.CMD_S2C_USER_TOKEN, new InvalidTokenResponseVo()).json()));
                    ctx.close();
                    return;
                }

                String userId = TokenHeper.getUserId(accessToken);
                if (userId == null || userId.isEmpty()) {
                    logger.debug("invalid token");
                    CmdVo cmdVo = new CmdVo(Command.CMD_S2C_USER_TOKEN, new InvalidTokenResponseVo());
                    logger.debug(cmdVo.json());
                    logger.debug(cmdVo.getBody().json());
                    ctx.channel().writeAndFlush(new TextWebSocketFrame(cmdVo.json()));
                    ctx.close();
                    return;
                }
                ChannelHub.addChannel(userId, ctx.channel());
            } else {
                HttpResponse response = new DefaultFullHttpResponse(request.protocolVersion(), HttpResponseStatus.OK, Unpooled.wrappedBuffer("SmartGate Websocket Server".getBytes()));
                response.headers().set(HttpHeaderNames.CONTENT_TYPE, "text/html; charset=UTF-8");
                ctx.writeAndFlush(response);
                ctx.close();
            }
        } else {      
            if (request.uri().startsWith("/notify?")) {
                Map<String, List<String>> params = splitQuery(new URI(request.uri()));
                NotificationVO notificationVO = new NotificationVO();
                notificationVO.setId(Long.valueOf(params.get("id").get(0)));
                notificationVO.setType(Integer.valueOf(params.get("type").get(0)));
                ChannelSession cs = ChannelHub.getChannelSession(params.get("user_id").get(0));
                ResponseVo responseVo = null;
                if (cs != null) {
                    cs.setUserData(notificationVO);
                    cs.writeAndFlush(new CmdVo(Command.CMD_S2C_NOTIFICATION, new ResponseVo(notificationVO)));
                    responseVo = new SuccessResponseVo();
                    logger.info("notify user {}", cs.getUserId());
                } else {
                    responseVo = new ResponseVo(1004, "用户不存在或未在线");
                    logger.warn(responseVo.json());
                }
                HttpResponse response = new DefaultFullHttpResponse(request.protocolVersion(), HttpResponseStatus.OK, Unpooled.wrappedBuffer(responseVo.jsonBytes()));
                response.headers().set(HttpHeaderNames.CONTENT_TYPE, "application/json; charset=UTF-8");
                ctx.writeAndFlush(response);
            }
            
            ctx.close();
        }
    }

    // 三者与：1.GET? 2.Upgrade头 包含websocket字符串? 3.Connection头 包含 Upgrade字符串?
    private boolean isWebSocketUpgrade(FullHttpRequest req) {
        HttpHeaders headers = req.headers();
        String upgradeHeader = headers.get(HttpHeaderNames.UPGRADE);
        String connHeader = headers.get(HttpHeaderNames.CONNECTION);
        return req.method().equals(HttpMethod.GET) && upgradeHeader != null && upgradeHeader.contains(WEBSOCKET_UPGRADE)
                && connHeader != null && connHeader.contains(WEBSOCKET_CONNECTION);
    }
    
    public static Map<String, List<String>> splitQuery(URI url) throws UnsupportedEncodingException {
        final Map<String, List<String>> query_pairs = new LinkedHashMap<String, List<String>>();
        final String[] pairs = url.getQuery().split("&");
        for (String pair : pairs) {
          final int idx = pair.indexOf("=");
          final String key = idx > 0 ? URLDecoder.decode(pair.substring(0, idx), "UTF-8") : pair;
          if (!query_pairs.containsKey(key)) {
            query_pairs.put(key, new LinkedList<String>());
          }
          final String value = idx > 0 && pair.length() > idx + 1 ? URLDecoder.decode(pair.substring(idx + 1), "UTF-8") : null;
          query_pairs.get(key).add(value);
        }
        return query_pairs;
      }

}
