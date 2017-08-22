package com.youstars.sgsocket;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.youstars.sgsocket.vo.CmdVo;
import com.youstars.sgsocket.vo.NotificationConsumeVo;
import com.youstars.sgsocket.vo.NotificationVO;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.handler.codec.http.websocketx.CloseWebSocketFrame;
import io.netty.handler.codec.http.websocketx.PingWebSocketFrame;
import io.netty.handler.codec.http.websocketx.PongWebSocketFrame;
import io.netty.handler.codec.http.websocketx.TextWebSocketFrame;
import io.netty.handler.codec.http.websocketx.WebSocketFrame;

public class WebSocketFrameHandler extends SimpleChannelInboundHandler<WebSocketFrame> {

    private static Logger logger = LoggerFactory.getLogger(WebSocketFrameHandler.class);
    
    @Override
    protected void channelRead0(ChannelHandlerContext ctx, WebSocketFrame frame) throws Exception {

        // 判断是否关闭链路指令
        if (frame instanceof CloseWebSocketFrame) {
            // handshaker.close(ctx.channel(), (CloseWebSocketFrame) frame.retain());
            ChannelHub.removeChannel(ctx.channel());
            return;
        }
        // 判断是否Ping消息 -- ping/pong心跳包
        if (frame instanceof PingWebSocketFrame) {
            ctx.channel().write(new PongWebSocketFrame(frame.content().retain()));
            return;
        }
        // 本程序仅支持文本消息， 不支持二进制消息
        if (!(frame instanceof TextWebSocketFrame)) {
            throw new UnsupportedOperationException(
                    String.format("%s frame types not supported", frame.getClass().getName()));
        }

        // 返回应答消息 text文本帧
        String request = ((TextWebSocketFrame) frame).text();
        
        try {
            ObjectMapper mapper = new ObjectMapper();
            CmdVo cmdVo = mapper.readValue(request.getBytes(), CmdVo.class);
            if (cmdVo.getCmd() == Command.CMD_C2S_NOTIFICATION_CONSUME) {
                NotificationConsumeVo nvo = mapper.readValue(request.getBytes(), NotificationConsumeVo.class);
                ChannelSession cs = ChannelHub.getChannelSession(ctx.channel());
                if (cs != null) {
                    NotificationVO _nvo = (NotificationVO) cs.getUserData();
                    if (_nvo != null && _nvo.getId() == nvo.getId()) {
                        cs.setUserData(null);
                        logger.info("user consumed notification {}", nvo.getId());
                    }
                }
            } else {
                cmdVo.setCmd(Command.CMD_S2C_NOT_FOUND);
                logger.warn(cmdVo.json());
                ctx.channel().writeAndFlush(new TextWebSocketFrame(cmdVo.json()));
            }
        } catch (Exception e) {
            logger.error(e.getMessage());
            logger.debug("pong request message");
            ctx.channel().writeAndFlush(new TextWebSocketFrame("pong: " + request));
        }
    }

}
