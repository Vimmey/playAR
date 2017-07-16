package ar.swiggy.com.swiggyar;

import android.app.Activity;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.Log;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * Created by Anupam on 16/07/17.
 */

public class ListingActivity extends Activity {

    RecyclerView recyclerView;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.listing_activity);

        recyclerView = (RecyclerView) findViewById(R.id.listing_recycler_view);
        recyclerView.setLayoutManager(new LinearLayoutManager(this));


        APIInterface apiService = APIClient.getClient().create(APIInterface.class);

        Call<ListingTopHirerachy> call = apiService.getRest();
        call.enqueue(new Callback<ListingTopHirerachy>() {
            @Override
            public void onResponse(Call<ListingTopHirerachy> call, Response<ListingTopHirerachy> response) {
                ListingTopHirerachy listing = response.body();
                recyclerView.setAdapter(new ListingAdapter(listing, R.layout.listing_list_item, getApplicationContext()));
            }

            @Override
            public void onFailure(Call<ListingTopHirerachy> call, Throwable t) {
                // Log error here since request failed
                Log.e("", t.toString());
            }
        });
    }
}
