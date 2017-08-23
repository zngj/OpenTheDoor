/**
 * 
 */
package com.youstars.sgsocket.vo;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
 * @author carson
 *
 */
@JsonIgnoreProperties(ignoreUnknown=true)
public class CmdVo {

    protected int cmd;
    
    protected ResponseVo body;
    
    public CmdVo() {
        
    }
    
    public CmdVo(int cmd, ResponseVo body) {
        this.cmd = cmd;
        this.body = body;
    }
    
    public String json() throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.writeValueAsString(this);
    }

    public int getCmd() {
        return cmd;
    }

    public void setCmd(int cmd) {
        this.cmd = cmd;
    }

    public ResponseVo getBody() {
        return body;
    }

    public void setBody(ResponseVo body) {
        this.body = body;
    }

}
