package ar.swiggy.com.swiggyar;

import com.google.gson.annotations.SerializedName;

import java.io.Serializable;

/**
 * Created by Anupam on 15/07/17.
 */

class MenuModel implements Serializable {

    @SerializedName("item_name")
    private String mItemName;

    @SerializedName("item_price")
    private String mItemPrice;

    public MenuModel(String mItemName, String mItemPrice) {
        this.mItemName = mItemName;
        this.mItemPrice = mItemPrice;
    }

    String getItemName() {
        return mItemName;
    }

    String getItemPrice() {
        return mItemPrice;
    }

}
