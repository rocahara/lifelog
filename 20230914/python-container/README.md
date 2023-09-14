## やったこと
- dockerコンテナでflaskを起動し、ローカルからアクセスする。

## はまったこと
- サーバーは起動して想定通りのポートでLISTENしているのに、ローカルからアクセスするとページが見えない。
- app.runでhost="0.0.0.0"を指定するとアクセスできるようになった。
  - (参考URL)https://qiita.com/amuyikam/items/01a8c16e3ddbcc734a46