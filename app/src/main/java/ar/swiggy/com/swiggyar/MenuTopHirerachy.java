package ar.swiggy.com.swiggyar;

import com.google.gson.annotations.SerializedName;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by Anupam on 15/07/17.
 */

class MenuTopHirerachy implements Serializable {

    @SerializedName("menu_details")
    private List<MenuModel> menuDetails = new ArrayList<>();

    List<MenuModel> getMenuDetails() {
        return menuDetails;
    }
}
