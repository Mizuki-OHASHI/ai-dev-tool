# 2023-09-18 　 OVERVIEW

- Craft の publish_contents_templates を例に reference, model から controller, usecase, dao のコードをどれくらい生成ができるか試してみる
  - **目的** 課題の発見
  - **内容** 全体 / package ごと / 関数ごと に生成してみる
  - モデル
    - (chatGPT3.5)
    - gpt-s.5-turbo
- 自動でデバッグするコードの実装
  - **目標** ロジックの補強
  - **内容** (以下を自動的に行う)
    - 生成されたコードを実行
    - エラーをプロンプトに加える
    - これを繰り返す
  - モデル
    - gpt-s.5-turbo
  - **課題**
    - テストケースの自動生成
- Prompt Engineering: https://www.promptingguide.ai/jp
