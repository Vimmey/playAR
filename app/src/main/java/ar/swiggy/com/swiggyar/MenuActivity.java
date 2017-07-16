package ar.swiggy.com.swiggyar;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v7.widget.LinearLayoutManager;
import android.support.v7.widget.RecyclerView;
import android.util.Log;
import android.view.View;
import android.widget.RelativeLayout;

import retrofit2.Call;
import retrofit2.Callback;
import retrofit2.Response;

/**
 * Created by Anupam on 15/07/17.
 */

public class MenuActivity extends Activity {

    private static final String TAG = "Menu";

    RecyclerView recyclerView;
    RelativeLayout cartCheckout;
    Response<MenuTopHirerachy> mResponse;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.menu_activity);

        cartCheckout = (RelativeLayout) findViewById(R.id.cartCheckout);
        recyclerView = (RecyclerView) findViewById(R.id.menu_recycler_view);
        recyclerView.setLayoutManager(new LinearLayoutManager(this));

        APIInterface apiService = APIClient.getClient().create(APIInterface.class);

        Call<MenuTopHirerachy> call = apiService.getMenuList();
        call.enqueue(new Callback<MenuTopHirerachy>() {
            @Override
            public void onResponse(Call<MenuTopHirerachy> call, Response<MenuTopHirerachy> response) {
                Log.d("check--", response.toString());
                mResponse = response;
                MenuTopHirerachy menu = response.body();
                recyclerView.setAdapter(new MenuAdapter(menu, R.layout.menu_list_item, getApplicationContext()));
            }

            @Override
            public void onFailure(Call<MenuTopHirerachy> call, Throwable t) {
                Log.e(TAG, t.toString());
            }
        });

        cartCheckout.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(MenuActivity.this, TrackActivity.class);
                startActivity(intent);
            }
        });
    }
}
