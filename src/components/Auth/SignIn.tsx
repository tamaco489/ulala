import { API_HOST, API_REQUEST_OPTIONS } from '@/config/env';
import axios from 'axios';
import { FirebaseError } from 'firebase/app';
import { getAuth, signInWithEmailAndPassword, getIdToken } from 'firebase/auth';
import { useRouter } from 'next/router';
import { FormEvent, useState } from 'react';

type SignInResponse = {
  uid: number;
};

const SignIn = () => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isToastOpen, setIsToastOpen] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<string>('');
  const { push } = useRouter();

  const handleSignIn = async (e: FormEvent<HTMLFormElement>) => {
    setIsLoading(true);
    e.preventDefault();
    try {
      const auth = getAuth();
      const userCredentials = await signInWithEmailAndPassword(
        auth,
        email,
        password
      );

      const loginUser = auth.currentUser;
      if (loginUser === null) {
        return console.log('not login');
      }
      const idToken = await getIdToken(loginUser, true);
      console.log('firebase token:', idToken);

      const signInResponse = await axios.post<SignInResponse>(
        `${API_HOST}/signin`,
        {
          firebase_token: idToken,
          email: email,
        },
        API_REQUEST_OPTIONS
      );
      console.log('signin response: ', signInResponse.data);

      setIsToastOpen(true);
      setEmail('');
      setPassword('');
    } catch (e: any) {
      setIsToastOpen(true);
      setErrorMessage('Firebase Authentication SignIn Failed');
      if (e instanceof FirebaseError) {
        console.log(`error code: ${e.code}, error message: ${e.message}`);
      }
    } finally {
      setIsLoading(false);
      setTimeout(() => {
        setIsToastOpen(false);
      }, 3000);
      push('/');
    }
  };

  return (
    <>
      <form className="mt-8" onSubmit={handleSignIn}>
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
            {isLoading ? 'Processing' : 'SignIn'}
          </button>

          {isToastOpen && (
            <div
              className={`fixed bottom-4 right-4 ${
                errorMessage ? 'bg-red-500' : 'bg-green-500'
              } text-white py-2 px-4 rounded`}
            >
              {errorMessage
                ? 'Email address or password is invalid'
                : 'Sign in Successfully'}
            </div>
          )}
        </div>
      </form>
    </>
  );
};

export default SignIn;
