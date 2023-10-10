## Auto Debug テスト

### 概要

**目的**: package autodebug の動作確認
**結果**: 概ね良好, ロジックの間違いを検出・訂正できた

### 内容

- 簡単な数学的な問題を与え, その答えを返す関数を生成する
- main 関数はあらかじめ与えておく
- main 関数でテストを実行し, そのログをフィードバックとして再び GPT にリクエストする

| Test No. | Problem                               | Results         | Remarks      | Directly |
| :------- | :------------------------------------ | :-------------- | :----------- | -------- |
| 1        | factorial                             | fail -> success | simple debug | 01,02,03 |
| 2-1      | cubiculum/Go/Assignment1              | success         |              | 04       |
| 2-2      | cubiculum/Go/Assignment2              | success         |              | 05       |
| 2-3      | cubiculum/Go/Assignment2 (poor hints) | fail -> success | modify query | 06,07    |
