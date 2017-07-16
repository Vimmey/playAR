package ar.swiggy.com.swiggyar;

import android.content.Intent;
import android.content.SharedPreferences;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.TableLayout;
import android.widget.TableRow;
import android.widget.TextView;
import android.widget.Toast;

/**
 * Created by Anupam on 16/07/17.
 */

public class Board extends AppCompatActivity {

    private int size;
    TableLayout mainBoard;
    TextView tv_turn;
    TextView player1Score, player2Score;
    char[][] board;
    char turn;
    int player1Wins, player2Wins;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_board);

        size = Integer.parseInt(getString(R.string.size_of_board));
        board = new char[size][size];
        mainBoard = (TableLayout) findViewById(R.id.mainBoard);
        tv_turn = (TextView) findViewById(R.id.turn);

        player1Score = (TextView) findViewById(R.id.player1Score);
        player2Score = (TextView) findViewById(R.id.player2Score);

        resetBoard();
        tv_turn.setText("Turn: Player 1");

        for (int i = 0; i < mainBoard.getChildCount(); i++) {
            TableRow row = (TableRow) mainBoard.getChildAt(i);
            for (int j = 0; j < row.getChildCount(); j++) {
                TextView tv = (TextView) row.getChildAt(j);
                tv.setText(R.string.none);
                tv.setOnClickListener(Move(i, j, tv));
            }
        }

        Button rstbtn = (Button) findViewById(R.id.reset);
        rstbtn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent current = getIntent();
                finish();
                startActivity(current);
            }
        });

        SharedPreferences prefs = getSharedPreferences("Score", MODE_PRIVATE);
        String player1ScoreStr = prefs.getString("player1", "");
        String player2ScoreStr = prefs.getString("player2", "");

        player1Score.setText(player1ScoreStr);
        player2Score.setText(player2ScoreStr);

    }

    protected void resetBoard() {
        turn = 'X';
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                board[i][j] = ' ';
            }
        }
    }

    protected int gameStatus() {

        //0 Continue
        //1 X Wins
        //2 O Wins
        //-1 Draw

        int rowX = 0, colX = 0, rowO = 0, colO = 0;
        for (int i = 0; i < size; i++) {
            if (check_Row_Equality(i, 'X'))
                return 1;
            if (check_Column_Equality(i, 'X'))
                return 1;
            if (check_Row_Equality(i, 'O'))
                return 2;
            if (check_Column_Equality(i, 'O'))
                return 2;
            if (check_Diagonal('X'))
                return 1;
            if (check_Diagonal('O'))
                return 2;
        }

        boolean boardFull = true;
        for (int i = 0; i < size; i++) {
            for (int j = 0; j < size; j++) {
                if (board[i][j] == ' ')
                    boardFull = false;
            }
        }
        if (boardFull)
            return -1;
        else return 0;
    }

    protected boolean check_Diagonal(char player) {
        int count_Equal1 = 0, count_Equal2 = 0;
        for (int i = 0; i < size; i++)
            if (board[i][i] == player)
                count_Equal1++;
        for (int i = 0; i < size; i++)
            if (board[i][size - 1 - i] == player)
                count_Equal2++;
        if (count_Equal1 == size || count_Equal2 == size)
            return true;
        else return false;
    }

    protected boolean check_Row_Equality(int r, char player) {
        int count_Equal = 0;
        for (int i = 0; i < size; i++) {
            if (board[r][i] == player)
                count_Equal++;
        }

        if (count_Equal == size)
            return true;
        else
            return false;
    }

    protected boolean check_Column_Equality(int c, char player) {
        int count_Equal = 0;
        for (int i = 0; i < size; i++) {
            if (board[i][c] == player)
                count_Equal++;
        }

        if (count_Equal == size)
            return true;
        else
            return false;
    }

    protected boolean Cell_Set(int r, int c) {
        return !(board[r][c] == ' ');
    }

    protected void stopMatch() {
        for (int i = 0; i < mainBoard.getChildCount(); i++) {
            TableRow row = (TableRow) mainBoard.getChildAt(i);
            for (int j = 0; j < row.getChildCount(); j++) {
                TextView tv = (TextView) row.getChildAt(j);
                tv.setOnClickListener(null);
            }
        }
    }

    View.OnClickListener Move(final int r, final int c, final TextView tv) {

        return new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                if (!Cell_Set(r, c)) {
                    board[r][c] = turn;
                    if (turn == 'X') {
                        tv.setText(R.string.X);
                        turn = 'O';
                    } else if (turn == 'O') {
                        tv.setText(R.string.O);
                        turn = 'X';
                    }
                    if (gameStatus() == 0) {
                        if (turn == 'X')
                            tv_turn.setText("Turn: Player 1");
                        else
                            tv_turn.setText("Turn: Player 2");
                    } else if (gameStatus() == -1) {
                        tv_turn.setText("Game: Draw");
                        stopMatch();
                    } else {
                        if (turn == 'X') {
                            tv_turn.setText("Player 1 Loses!");
                            player2Wins++;
                            player2Score.setText(String.valueOf(player2Wins));
                        } else {
                            tv_turn.setText("Player 2 Loses!");
                            player1Wins++;
                            player1Score.setText(String.valueOf(player1Wins));
                        }
                        saveScore();
                        stopMatch();
                    }
                } else {
                    Toast.makeText(getApplicationContext(), "Please choose a cell which is not already occupied", Toast.LENGTH_LONG).show();
                }
            }
        };
    }

    private void saveScore() {
        SharedPreferences.Editor editor = getSharedPreferences("Score", MODE_PRIVATE).edit();
        if (turn == 'X') {
            editor.putString("player2", String.valueOf(player2Wins));
        } else {
            editor.putString("player1", String.valueOf(player1Wins));
        }
        editor.apply();
    }
}
