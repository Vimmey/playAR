package ar.swiggy.com.swiggyar;

import com.google.gson.annotations.SerializedName;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by Anupam on 15/07/17.
 */

public class MenuTopHirerachy implements Serializable {

    @SerializedName("menu_details")
    private List<MenuModel> menuDetails = new ArrayList<>();

    public List<MenuModel> getMenuDetails() {
        return menuDetails;
    }

    public void setMenuDetails(List<MenuModel> menuDetails) {
        this.menuDetails = menuDetails;
    }
}
