export default function Footer() {
  const year = new Date().getFullYear();

  return (
    <div className="fixed bottom-0 w-full bg-gray-800 text-white text-center p-8">
      <h2>© hoge Inc 2022-{year}</h2>
    </div>
  );
}
