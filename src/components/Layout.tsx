import Header from './Header';
import Footer from './Footer';

export default function Layout({ children }: any) {
  return (
    <>
      {/* <h1>どのページでも表示する共通のコンポーネント</h1> */}
      <Header />
      <div className="container mx-auto px-5 w-2/3 lg:w-2/5 xl:w-1/3">
        {children}
      </div>
      <Footer />
    </>
  );
}
