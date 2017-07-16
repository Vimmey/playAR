package ar.swiggy.com.swiggyar;

import com.google.gson.annotations.SerializedName;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

/**
 * Created by Anupam on 16/07/17.
 */

public class ListingTopHirerachy implements Serializable {

    @SerializedName("restaurants")
    private List<ListingModel> restDetails = new ArrayList<>();

    public List<ListingModel> getRest() {
        return restDetails;
    }
}
