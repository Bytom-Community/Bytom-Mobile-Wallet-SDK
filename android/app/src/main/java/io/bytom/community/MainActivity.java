package io.bytom.community;

import android.app.Activity;
import android.os.Bundle;
import android.os.Environment;
import android.util.Log;
import android.widget.TextView;

import wallet.Wallet;

public class MainActivity extends Activity {

    private TextView mTextView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        mTextView = (TextView) findViewById(R.id.mytextview);

//        String keystorePath = getFilesDir().toString();
        String keystorePath = getExternalFilesDir(Environment.DIRECTORY_DOWNLOADS).toString();
        Log.d("dataDir", "Environment.getDataDirectory()=:" + keystorePath);
//      Call Go function.
        String publicKey = Wallet.createWallet(keystorePath, "marshall", "123456");
        mTextView.setText(publicKey);
    }
}
