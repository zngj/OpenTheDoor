/**
 * 
 */
package com.youstars.sgsocket.vo;

/**
 * @author carson
 *
 */
public class InvalidTokenResponseVo extends ResponseVo {
    public InvalidTokenResponseVo() {
        this.code = 1000;
        this.msg = "登录失效";
    }
}
