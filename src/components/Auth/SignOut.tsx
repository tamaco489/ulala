import { useState } from 'react';
import { useRouter } from 'next/router';
import { useAuthContext } from '@/feature/auth/provider/AuthProvider';
import { FirebaseError } from 'firebase/app';
import { getAuth, getIdToken, signOut } from 'firebase/auth';
import axios from 'axios';
import { API_HOST, API_REQUEST_OPTIONS } from '@/config/env';

const SignOut = () => {
  const { user } = useAuthContext();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isToastOpen, setIsToastOpen] = useState<boolean>(false);
  const { push } = useRouter();

  const handleSignOut = async () => {
    setIsLoading(true);
    try {
      const auth = getAuth();
      const loginUser = auth.currentUser;
      if (loginUser === null) {
        return console.log('Invalid firebase user');
      }
      const idToken = await getIdToken(loginUser, true);
      await axios.post(
        `${API_HOST}/signout`,
        {
          firebase_token: idToken,
        },
        API_REQUEST_OPTIONS
      );

      await signOut(auth);
      setIsToastOpen(true);
      push('/signin');
    } catch (e: any) {
      if (e instanceof FirebaseError) {
        console.log(`error code: ${e.code}, error message: ${e.message}`);
      }
    } finally {
      setIsLoading(false);
      setTimeout(() => {
        setIsToastOpen(false);
      }, 3000);
      push('/signin');
    }
  };

  return (
    <>
      {user ? (
        <button
          className="block bg-gray-500 text-white py-2 px-4 hover:bg-yellow-400 rounded transition-all duration-300"
          onClick={handleSignOut}
          disabled={isLoading}
        >
          {isLoading ? 'SignOut Processing' : 'SignOut'}
        </button>
      ) : (
        <p>Please SignIn</p>
      )}

      {isToastOpen ? (
        <p className="fixed bottom-4 right-4 bg-green-500 text-white py-2 px-4 rounded">
          Success signout
        </p>
      ) : (
        <></>
      )}
    </>
  );
};

export default SignOut;
