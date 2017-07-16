package ar.swiggy.com.swiggyar;

import retrofit2.Call;
import retrofit2.http.GET;
import retrofit2.http.Query;

/**
 * Created by Anupam on 15/07/17.
 */

interface APIInterface {

    @GET("menu/363222332")
    Call<MenuTopHirerachy> getMenuList();

    @GET("/v2/restaurants")
    Call<ListingTopHirerachy> getRest();
}
