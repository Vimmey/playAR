package ar.swiggy.com.swiggyar;

import com.google.gson.annotations.SerializedName;

import java.util.List;

/**
 * Created by Anupam on 15/07/17.
 */

public class MenuResponse {

    @SerializedName("results")
    private List<MenuModel> results;

    public List<MenuModel> getResults() {
        return results;
    }
}
