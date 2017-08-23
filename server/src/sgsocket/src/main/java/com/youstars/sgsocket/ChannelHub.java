package com.youstars.sgsocket;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.youstars.sgsocket.vo.CmdVo;

import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.util.Attribute;
import io.netty.util.AttributeKey;

public class ChannelHub {

    private static Logger logger = LoggerFactory.getLogger(HttpRequestHandler.class);

    private static Map<String, ChannelSession> channels = new ConcurrentHashMap<String, ChannelSession>();

    public static void addChannel(String userId, Channel ch) {
        ChannelSession cs = channels.get(userId);
        if (cs != null) {
            cs.getChannel().close();
            logger.info("closed user {} channel {}", userId, cs.getChannel().id());
        }
        AttributeKey<String> attrKey = AttributeKey.valueOf("USER_ID");
        Attribute<String> attr = ch.attr(attrKey);
        attr.set(userId);
        channels.put(userId, new ChannelSession(userId, ch));
        logger.info("added user {} new channel {}", userId, ch.id());
        logger.info("channel count: {}", channels.size());
    }
    
    public static String getUserId(Channel ch) {
        AttributeKey<String> attrKey = AttributeKey.valueOf("USER_ID");
        Attribute<String> attr = ch.attr(attrKey);
        if (attr != null) {
            return attr.get();
        }
        return null;
    }
    
    public static ChannelSession getChannelSession(String userId) {
        return channels.get(userId);
    }
    
    public static ChannelSession getChannelSession(Channel ch) {
        return channels.get(getUserId(ch));
    }

    public static void removeChannel(Channel ch) {
        String userId = getUserId(ch);
        if (userId != null) {
            removeChannel(userId);
        }
    }

    public static void removeChannel(String userId) {
        ChannelSession cs = channels.remove(userId);
        if (cs != null) {
            cs.getChannel().close();
            logger.info("user {} channel {} was closed", userId, cs.getChannel().id());
            logger.info("channel count: {}", channels.size());
        }
    }

    public static ChannelFuture writeAndFlush(String userId, CmdVo vo) throws Exception {
        ChannelSession cs = channels.get(userId);
        if (cs == null) {
            logger.info("user {} not connnected", userId);
            return null;
        }
        return cs.writeAndFlush(vo);
    }
}
