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
 * Created by Anupam on 16/07/17.
 */

public class ListingAdapter extends RecyclerView.Adapter<ListingAdapter.RestViewHolder> {

    private List<ListingModel> listingModels;
    private int rowLayout;
    private Context context;

    public static class RestViewHolder extends RecyclerView.ViewHolder {
        LinearLayout restLayout;
        TextView mTitle;
        TextView mDesc;
        TextView mrating;
        TextView mCFT;

        public RestViewHolder(View v) {
            super(v);
            restLayout = (LinearLayout) v.findViewById(R.id.restListItem);
            mTitle = (TextView) v.findViewById(R.id.restName);
            mrating = (TextView) v.findViewById(R.id.restRating);
            mCFT = (TextView) v.findViewById(R.id.restCFT);
        }
    }

    public ListingAdapter(ListingTopHirerachy listingTopHirerachy, int rowLayout, Context context) {
        this.listingModels = listingTopHirerachy.getRest();
        this.rowLayout = rowLayout;
        this.context = context;
    }

    @Override
    public ListingAdapter.RestViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        View view = LayoutInflater.from(parent.getContext()).inflate(rowLayout, parent, false);
        return new RestViewHolder(view);
    }


    @Override
    public void onBindViewHolder(RestViewHolder holder, final int position) {
        holder.mTitle.setText(listingModels.get(position).getRestName());
        holder.mrating.setText("Rating " + listingModels.get(position).getRestRating());
        holder.mCFT.setText("₹" + listingModels.get(position).getRestCFT());
    }

    @Override
    public int getItemCount() {
        return listingModels.size();
    }
}
