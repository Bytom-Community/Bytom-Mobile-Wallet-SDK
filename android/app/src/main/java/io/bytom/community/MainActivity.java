package io.bytom.community;

import android.app.Activity;
import android.os.Bundle;
import android.util.Log;
import android.widget.TextView;

import org.json.JSONException;
import org.json.JSONObject;

import wallet.Wallet;

public class MainActivity extends Activity {

    @Override
    public void onBackPressed() {
        super.onBackPressed();
        android.os.Process.killProcess(android.os.Process.myPid());
        System.exit(0);
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        TextView keyTextView = (TextView) findViewById(R.id.key_textview);
        TextView accountTextView = (TextView) findViewById(R.id.account_textview);
        TextView addressTextView = (TextView) findViewById(R.id.address_textview);

        String storagePath = getFilesDir().toString();
        Log.d("storagePath", storagePath);

        Wallet.initWallet(storagePath);
        String keyResult = Wallet.createKey("marshall", "123456");
        Log.d("keyResult", keyResult);
        keyTextView.setText(keyResult);

        String xpub = "";
        try {
            JSONObject keyObject = new JSONObject(keyResult);
            String result = keyObject.getString("status");
            if (result.equals("success")) {
                xpub = keyObject.getJSONObject("data").getString("xpub");
            }
        } catch (JSONException e) {
            e.printStackTrace();
        }

        String accountResult = Wallet.createAccount("marshall", 1, xpub);
        Log.d("accountResult", accountResult);
        accountTextView.setText(accountResult);

        String addressResult = Wallet.createAccountReceiver("", "marshall");
        Log.d("address", addressResult);
        addressTextView.setText(addressResult);
    }
}
