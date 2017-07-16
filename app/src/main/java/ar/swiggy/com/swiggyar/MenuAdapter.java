package ar.swiggy.com.swiggyar;

import android.content.Context;
import android.support.v7.widget.RecyclerView;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.LinearLayout;
import android.widget.TextView;

import java.util.List;

/**
 * Created by Anupam on 15/07/17.
 */

public class MenuAdapter extends RecyclerView.Adapter<MenuAdapter.MenuViewHolder> {

    private List<MenuModel> menuModels;
    private int rowLayout;
    private Context context;

    public static class MenuViewHolder extends RecyclerView.ViewHolder {
        LinearLayout menuLayout;
        TextView menuName;
        TextView menuPrice;

        public MenuViewHolder(View v) {
            super(v);
            menuLayout = (LinearLayout) v.findViewById(R.id.menu_layout);
            menuName = (TextView) v.findViewById(R.id.item_name);
            menuPrice = (TextView) v.findViewById(R.id.item_price);
        }
    }

    MenuAdapter(MenuTopHirerachy menuModels, int rowLayout, Context context) {
        this.menuModels = menuModels.getMenuDetails();
        this.rowLayout = rowLayout;
        this.context = context;
    }

    @Override
    public MenuAdapter.MenuViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(rowLayout, parent, false);
        return new MenuViewHolder(view);
    }


    @Override
    public void onBindViewHolder(MenuViewHolder holder, final int position) {
        holder.menuName.setText(menuModels.get(position).getItemName());
        holder.menuPrice.setText("₹" + menuModels.get(position).getItemPrice());
    }

    @Override
    public int getItemCount() {
        return menuModels.size();
    }
}
