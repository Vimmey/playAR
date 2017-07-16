package ar.swiggy.com.swiggyar;

import android.Manifest;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.os.Bundle;
import android.support.annotation.Nullable;
import android.support.v4.app.ActivityCompat;
import android.support.v4.content.ContextCompat;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.widget.Toast;

import com.wikitude.architect.ArchitectJavaScriptInterfaceListener;
import com.wikitude.architect.ArchitectStartupConfiguration;
import com.wikitude.architect.ArchitectView;

import org.json.JSONException;
import org.json.JSONObject;

import java.io.IOException;

public class MainActivity extends AppCompatActivity implements AccelerometerListener {

    private static final int WIKITUDE_PERMISSIONS_REQUEST_CAMERA = 1;
    private static final int WIKITUDE_PERMISSIONS_REQUEST_GPS = 2;

    private static final String TAG = "MainActivity";

    ArchitectView architectView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        cameraPermissionRequest();

        this.architectView = (ArchitectView) this.findViewById(R.id.architectView);
        final ArchitectStartupConfiguration config = new ArchitectStartupConfiguration();
        config.setLicenseKey(StringConstants.licenseKey);
        this.architectView.onCreate(config);
    }

    @Override
    protected void onPostCreate(@Nullable Bundle savedInstanceState) {
        super.onPostCreate(savedInstanceState);
        this.architectView.onPostCreate();
        architectView.addArchitectJavaScriptInterfaceListener(new ArchitectJavaScriptInterfaceListener() {
            @Override
            public void onJSONObjectReceived(JSONObject jsonObject) {
                try {
                    switch (jsonObject.getString("name")) {
                        case "markerselected":
                            final Intent poiDetailIntent = new Intent(MainActivity.this, MenuActivity.class);
                            startActivity(poiDetailIntent);
                            break;
                    }
                } catch (JSONException e) {
                    Log.e(TAG, "onJSONObjectReceived: ", e);
                }
            }
        });
    }

    @Override
    protected void onPostResume() {
        super.onPostResume();
        try {
            architectView.setLocation(17.934847, 75.616123, 5);
            this.architectView.load("index.html");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    private void cameraPermissionRequest() {
        if (ContextCompat.checkSelfPermission(this, android.Manifest.permission.CAMERA) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(this, new String[]{android.Manifest.permission.CAMERA}, WIKITUDE_PERMISSIONS_REQUEST_CAMERA);
        } else {
            if (ContextCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
                ActivityCompat.requestPermissions(this, new String[]{Manifest.permission.ACCESS_FINE_LOCATION}, WIKITUDE_PERMISSIONS_REQUEST_GPS);
            }
        }
    }

    @Override
    public void onRequestPermissionsResult(int requestCode, String[] permissions, int[] grantResults) {
        switch (requestCode) {
            case WIKITUDE_PERMISSIONS_REQUEST_CAMERA: {
                if (grantResults.length > 0 && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                    if (ContextCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
                        ActivityCompat.requestPermissions(this, new String[]{Manifest.permission.ACCESS_FINE_LOCATION}, WIKITUDE_PERMISSIONS_REQUEST_GPS);
                    }
                } else {
                    Toast.makeText(this, "Sorry, augmented reality doesn't work without reality.\n\nPlease grant camera permission.", Toast.LENGTH_LONG).show();
                }
                return;
            }
            case WIKITUDE_PERMISSIONS_REQUEST_GPS: {
                if (grantResults.length > 0 && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                    Log.d(TAG, "onRequestPermissionsResult: ");
                } else {
                    Toast.makeText(this, "Sorry, this example requires access to your location in order to work properly.\n\nPlease grant location permission.", Toast.LENGTH_SHORT).show();
                }
                return;
            }
        }
    }

    public void onResume() {
        super.onResume();
        architectView.onResume();
        if (AccelerometerManager.isSupported(this)) {
            //Start Accelerometer Listening
            AccelerometerManager.startListening(this);
        }
    }

    protected void onPause() {
        super.onPause();
        if (architectView != null) {
            architectView.onPause();
        }
    }

    protected void onStop() {
        super.onStop();
        if (AccelerometerManager.isListening()) {
            //Start Accelerometer Listening
            AccelerometerManager.stopListening();
        }
    }

    protected void onDestroy() {
        super.onDestroy();
        if (architectView != null) {
            architectView.onDestroy();
        }
    }

    public void onLowMemory() {
        super.onLowMemory();
        if (architectView != null) {
            architectView.onLowMemory();
        }
    }

    @Override
    public void onAccelerationChanged(float x, float y, float z) {

    }

    @Override
    public void onShake(float force) {
//        Intent intent = new Intent(MainActivity.this, MenuActivity.class);
//        startActivity(intent);
//        finish();
    }
}
