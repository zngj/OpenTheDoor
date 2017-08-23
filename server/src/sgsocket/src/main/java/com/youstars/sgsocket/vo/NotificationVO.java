/**
 * 
 */
package com.youstars.sgsocket.vo;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonInclude.Include;

/**
 * @author carson
 *
 */
@JsonInclude(Include.NON_DEFAULT)
public class NotificationVO extends DataVo {
    
    private long id;
    
    private int type;

    public long getId() {
        return id;
    }

    public void setId(long id) {
        this.id = id;
    }

    public int getType() {
        return type;
    }

    public void setType(int type) {
        this.type = type;
    }

}
