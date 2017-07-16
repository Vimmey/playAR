package ar.swiggy.com.swiggyar;

import com.google.gson.annotations.SerializedName;

import java.io.Serializable;

/**
 * Created by Anupam on 16/07/17.
 */

public class ListingModel implements Serializable {

    @SerializedName("title")
    private String mTitle;

    @SerializedName("description")
    private String mDesc;

    @SerializedName("rating")
    private String mRating;

    @SerializedName("costForTwo")
    private String mCFT;

    public ListingModel(String mTitle, String mDesc, String mRating, String mCFT) {
        this.mTitle = mTitle;
        this.mDesc = mDesc;
        this.mRating = mRating;
        this.mCFT = mCFT;
    }

    String getRestName() {
        return mTitle;
    }

    String getRestDesc() {
        return mDesc;
    }

    String getRestRating() {
        return mRating;
    }

    String getRestCFT() {
        return mCFT;
    }
}
