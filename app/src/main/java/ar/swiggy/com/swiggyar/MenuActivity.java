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
    private final static String restID = "363222332";

    RecyclerView recyclerView;
    RelativeLayout cartCheckout;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.menu_activity);

        cartCheckout = (RelativeLayout) findViewById(R.id.cartCheckout);
        recyclerView = (RecyclerView) findViewById(R.id.menu_recycler_view);
        recyclerView.setLayoutManager(new LinearLayoutManager(this));

        APIInterface apiService = APIClient.getClient().create(APIInterface.class);

        Call<MenuTopHirerachy> call = apiService.getMenuList(restID);
        call.enqueue(new Callback<MenuTopHirerachy>() {
            @Override
            public void onResponse(Call<MenuTopHirerachy> call, Response<MenuTopHirerachy> response) {
                MenuTopHirerachy menu = response.body();
                recyclerView.setAdapter(new MenuAdapter(menu, R.layout.menu_list_item, getApplicationContext()));
                Log.d(TAG, "Number of Item received: " + menu.getMenuDetails().size());
            }

            @Override
            public void onFailure(Call<MenuTopHirerachy> call, Throwable t) {
                // Log error here since request failed
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
