package ar.swiggy.com.swiggyar;

import android.app.Activity;
import android.content.Intent;
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

public class ListingActivity extends Activity implements AccelerometerListener {

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

    @Override
    public void onAccelerationChanged(float x, float y, float z) {

    }

    @Override
    public void onShake(float force) {
        Intent intent = new Intent(ListingActivity.this, MainActivity.class);
        startActivity(intent);
        finish();
    }

    @Override
    public void onResume() {
        super.onResume();
        //Check device supported Accelerometer senssor or not
        if (AccelerometerManager.isSupported(this)) {
            //Start Accelerometer Listening
            AccelerometerManager.startListening(this);
        }
    }

    @Override
    public void onStop() {
        super.onStop();
        //Check device supported Accelerometer sensor or not
        if (AccelerometerManager.isListening()) {
            //Start Accelerometer Listening
            AccelerometerManager.stopListening();
        }
    }

    @Override
    public void onDestroy() {
        super.onDestroy();
        //Check device supported Accelerometer senssor or not
        if (AccelerometerManager.isListening()) {
            //Start Accelerometer Listening
            AccelerometerManager.stopListening();
        }

    }
}
