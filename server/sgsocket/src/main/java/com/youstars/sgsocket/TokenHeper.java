/**
 * 
 */
package com.youstars.sgsocket;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import redis.clients.jedis.Jedis;

/**
 * @author Carson
 *
 */
public class TokenHeper {
    
    private static Logger logger = LoggerFactory.getLogger(TokenHeper.class);
    
    public static String getUserId(String accessToken) {
        Jedis jedis = new Jedis("localhost");    
        String userId = jedis.hget("access_token:" + accessToken, "userid");
        jedis.close();
        logger.debug("token: {}, user= {}", accessToken, userId);
        return userId;
    }
}
