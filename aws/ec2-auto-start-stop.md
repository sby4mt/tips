
## IAMポリシーの作成
EC2インスタンスの起動停止を実行できるポリシーを作成します。  
CloudWatchにログを出力したくない場合は、logsのポリシーを削除します。
```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "ec2:StartInstances",
        "ec2:StopInstances"
      ],
      "Resource": "*"
    }
  ]
}
```

## IAMロールの作成
Lambdaが実行できるようにIAM Roleを作成します。  
信頼されたエンティティはLambdaを選択し作成したポリシーをアタッチします。

## Lambda Functionの作成
|項目|設定|
|--|--|
|関数名|任意|
|ランタイム|python3.8|
|実行ロール|作成したIAMロール|
```
import boto3

def lambda_handler(event, context):
    region = event['Region']
    instances = event['Instances']
    ec2 = boto3.client('ec2', region_name=region)
    if event['Action'] == 'start':
        ec2.start_instances(InstanceIds=instances)
    elif event['Action'] == 'stop':
        ec2.stop_instances(InstanceIds=instances)
    else:
        return 1
    return 0
```
- region = event[‘Region’]はイベントに渡されたリージョンを定義します。
- instances = event[‘Instances’]はイベントに渡されたインスタンスを定義します。
- ec2.start_instances(InstanceIds=instances)はインスタンスを起動することを定義します。
- ec2.stop_instances(InstanceIds=instances)はインスタンスを停止することを定義します。  
ここのコードでは、特定のリージョン、インスタンスを入力することは不要です。

## CloudWatch Eventsで起動停止のイベントスケジュールを作成
イベント例
|項目|設定|
|--|--|
|イベントパターン|スケジュール|
|Cron式|0 10 * * ? *|
|ターゲット|Lambda関数|
|機能|先ほど作成した関数|
|ターゲット入力を設定|{"Action":"stop","Region":"us-east-2","Instances":["i-XXXXXXXX"]}|