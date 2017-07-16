package ar.swiggy.com.swiggyar;

/**
 * Created by Anupam on 16/07/17.
 */

public interface AccelerometerListener {

    public void onAccelerationChanged(float x, float y, float z);

    public void onShake(float force);

}
