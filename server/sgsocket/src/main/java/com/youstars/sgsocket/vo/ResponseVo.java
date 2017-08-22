/**
 * 
 */
package com.youstars.sgsocket.vo;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
 * @author kingsoft
 *
 */
public class ResponseVo {

    protected int code;
    
    @JsonInclude(Include.NON_EMPTY)
    protected String msg;
    
    @JsonInclude(Include.NON_NULL)
    protected DataVo data;

    public ResponseVo() {
        
    }
    
    public ResponseVo(int code) {
        this.code = code;
    }

    public ResponseVo(String msg) {
        this.msg = msg;
    }
    
    public ResponseVo(DataVo data) {
        this.data = data;
    }
    
    public ResponseVo(int code, String msg) {
        this.code = code;
        this.msg = msg;
    }
    
    public ResponseVo(int code, String msg, DataVo data) {
        this.code = code;
        this.msg = msg;
        this.data = data;
    }
    
    public String json() throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.writeValueAsString(this);
    }
    
    public byte[] jsonBytes() throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.writeValueAsBytes(this);
    }
    
    public int getCode() {
        return code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public DataVo getData() {
        return data;
    }

    public void setData(DataVo data) {
        this.data = data;
    }
    
    
}
