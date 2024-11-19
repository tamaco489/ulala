import '@/styles/globals.css';
import type { AppProps } from 'next/app';
import Layout from '@/components/Layout';
import { initializeFirebaseApp } from '@/lib/firebase/firebase';
import { AuthProvider } from '@/feature/auth/provider/AuthProvider';

initializeFirebaseApp();
export default function App({ Component, pageProps }: AppProps) {
  // console.log(getApp()); // TODO: 後に削除
  return (
    <AuthProvider>
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </AuthProvider>
  );
}
