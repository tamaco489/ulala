import { load } from 'ts-dotenv';
import {
  Bucket,
  GetObjectCommand,
  ListBucketsCommand,
  PutObjectCommand,
  S3Client,
} from '@aws-sdk/client-s3';
import { fromIni } from '@aws-sdk/credential-providers';

const env = load({
  BUCKET_NAME: String,
  AWS_PROFILE: String,
  S3_ENDPOINT: String,
});

class S3Helper {
  private readonly s3Client: S3Client;
  constructor() {
    this.s3Client = new S3Client({
      credentials: fromIni({ profile: env.AWS_PROFILE }),
      endpoint: env.S3_ENDPOINT,
    });
  }

  /* バケット一覧を取得 */
  async listBuckets(): Promise<Bucket[] | undefined> {
    try {
      const res = await this.s3Client.send(new ListBucketsCommand({}));
      console.log('Success', res.Buckets);
      return res.Buckets;
    } catch (err) {
      console.log('Error', err);
    }
  }

  /* オブジェクトの取得 */
  async getObject(key: string): Promise<string | undefined> {
    try {
      const bucketParams = {
        Bucket: env.BUCKET_NAME,
        Key: key,
      };
      const res = await this.s3Client.send(new GetObjectCommand(bucketParams));
      return (await res.Body?.transformToString()) ?? '';
    } catch (err) {
      console.log('Error', err);
    }
  }

  /* オブジェクトのアップロード */
  async uploadObject(body: string, data: string): Promise<void> {
    try {
      const bucketParams = {
        Bucket: env.BUCKET_NAME,
        Key: data,
        Body: body,
      };
      const res = await this.s3Client.send(new PutObjectCommand(bucketParams));
      console.log('Success', res);
    } catch (err) {
      console.log('Error', err);
    }
  }
}

export const s3Helper = new S3Helper();
