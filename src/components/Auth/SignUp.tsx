import { FirebaseError } from 'firebase/app';
import {
  createUserWithEmailAndPassword,
  getAuth,
  getIdToken,
  sendEmailVerification,
} from 'firebase/auth';
import { FormEvent, useState } from 'react';
import { API_HOST, API_REQUEST_OPTIONS } from '@/config/env';
import axios from 'axios';
import { useRouter } from 'next/router';

type SignUpResponse = {
  id: string;
};

const SignUp = () => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isToastOpen, setIsToastOpen] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<string>('');
  const { push } = useRouter();

  const handleSingUp = async (e: FormEvent<HTMLFormElement>) => {
    setIsLoading(true);
    e.preventDefault();
    try {
      const auth = getAuth();

      /* 戻り値がPromiseなため、async関数内でawaitする必要あり */
      const userCredentials = await createUserWithEmailAndPassword(
        auth,
        email,
        password
      );

      /* createUserWithEmailAndPasswordの戻り値のオブジェクトのuserを引数に指定することで、ユーザに紐づけられたメールアドレスに確認メールを送信する */
      await sendEmailVerification(userCredentials.user);

      /* ログイン済みユーザのトークンを取得 */
      const loginUser = auth.currentUser;
      if (loginUser === null) {
        return console.log('Invalid firebase user');
      }
      const idToken = await getIdToken(loginUser, true);

      /* バックエンドサーバとの通信 */
      const signUpResponse = await axios.post<SignUpResponse>(
        `${API_HOST}/signup`,
        {
          firebase_token: idToken,
          email: email,
        },
        API_REQUEST_OPTIONS
      );
      console.log('signup response: ', signUpResponse);

      /* トーストメッセージを表示 */
      setIsToastOpen(true);

      /* 成功後、各フォームの入力値をリセットする */
      setEmail('');
      setPassword('');
    } catch (e: any) {
      setIsToastOpen(true);
      setErrorMessage('Firebase Authentication SignUp Failed');
      if (e instanceof FirebaseError) {
        console.log(`error code: ${e.code}, error message: ${e.message}`);
      }
    } finally {
      /* 3秒後にトーストメッセージを非表示にする */
      setIsLoading(false);
      setTimeout(() => {
        setIsToastOpen(false);
      }, 3000);
      push('/');
    }
  };

  return (
    <>
      <form className="mt-8" onSubmit={handleSingUp}>
        <div className="grid gap-4">
          <div>
            <label className="block">E-Mail Address</label>
            <input
              type="email"
              name="email"
              className="border border-gray-300 px-4 py-2 w-full"
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
            />
          </div>

          <div>
            <label className="block">Password</label>
            <input
              type="password"
              name="password"
              className="border border-gray-300 px-4 py-2 w-full"
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
              }}
            />
          </div>
        </div>

        <div className="mt-8">
          <button
            type="submit"
            className={`py-2 px-4 rounded ${
              isLoading ? 'bg-gray-400' : 'bg-blue-500'
            } text-white`}
            disabled={isLoading}
          >
            {isLoading ? 'Processing' : 'Create your Account'}
          </button>

          {isToastOpen && (
            <div
              className={`fixed bottom-4 right-4 ${
                errorMessage ? 'bg-red-500' : 'bg-green-500'
              } text-white py-2 px-4 rounded`}
            >
              {errorMessage
                ? 'Email address or password is invalid'
                : 'Send Your E-Mail Address'}
            </div>
          )}
        </div>
      </form>
    </>
  );
};

export default SignUp;
