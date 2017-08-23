/**
 * 
 */
package com.youstars.sgsocket;

import com.youstars.sgsocket.vo.CmdVo;

import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.handler.codec.http.websocketx.TextWebSocketFrame;

/**
 * @author carson
 *
 */
public class ChannelSession {

    private String userId;
    private Channel channel;
    private Object userData;
    private Object userData2;
    private Object userData3;
    //private Collection<Object> objects;
    
    public ChannelSession(String userId, Channel channel) {
        this.userId = userId;
        this.channel = channel;
    }

    public ChannelFuture writeAndFlush(CmdVo vo) throws Exception {
        return channel.writeAndFlush(new TextWebSocketFrame(vo.json()));
    }
    
    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }

    public Channel getChannel() {
        return channel;
    }

    public void setChannel(Channel channel) {
        this.channel = channel;
    }

    public Object getUserData() {
        return userData;
    }

    public void setUserData(Object userData) {
        this.userData = userData;
    }

    public Object getUserData2() {
        return userData2;
    }

    public void setUserData2(Object userData2) {
        this.userData2 = userData2;
    }

    public Object getUserData3() {
        return userData3;
    }

    public void setUserData3(Object userData3) {
        this.userData3 = userData3;
    }
    
}
