import { AuthGuard } from '@/feature/auth/secure/AuthGuard';
import Link from 'next/link';

export default function TopPage() {
  return (
    <AuthGuard>
      <div className="text-center my-10 mx-5 ">
        <div className="p-5 border border-gray-500 bg-gray-500 text-white">
          <h2>Firebase With Next.js Start</h2>
        </div>

        <div className="my-5">
          <Link href={'/movies/categories'}>
            <p>動画一覧ページへ</p>
          </Link>
        </div>
      </div>
    </AuthGuard>
  );
}
