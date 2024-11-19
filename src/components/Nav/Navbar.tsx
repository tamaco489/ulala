import { useAuthContext } from '@/feature/auth/provider/AuthProvider';
import Link from 'next/link';
import React from 'react';
import SignOut from '../Auth/SignOut';

export default function Navbar() {
  const { user } = useAuthContext();

  return (
    <div className="container mx-auto lg:px-2 px-5 lg:w-2/5">
      <div className="container flex items-center justify-between">
        <Link href="/" className="text-2xl font-medium">
          Common Navigation Bar
        </Link>
        <div>
          <ul className="flex items-center text-sm py-4">
            <li>
              <Link
                href="/"
                className="block bg-gray-500 mx-1 px-4 py-2 text-white rounded hover:bg-yellow-400 transition-all duration-300"
              >
                Top
              </Link>
            </li>
            {user ? (
              <li>
                <SignOut />
              </li>
            ) : (
              <>
                <li>
                  <Link
                    href="/signup"
                    className="block bg-gray-500 mx-1 px-4 py-2 text-white rounded hover:bg-yellow-400 transition-all duration-300"
                  >
                    SignUp
                  </Link>
                </li>
                <li>
                  <Link
                    href="/signin"
                    className="block bg-gray-500 mx-1 px-4 py-2 text-white rounded hover:bg-yellow-400 transition-all duration-300"
                  >
                    SignIn
                  </Link>
                </li>
              </>
            )}
            <li>
              <Link
                href="https://nextjs.org/docs"
                className="block bg-gray-500 mx-1 px-4 py-2 text-white rounded hover:bg-yellow-400 transition-all duration-300"
              >
                Next.js
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}
