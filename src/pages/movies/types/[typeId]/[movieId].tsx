import { useRouter } from 'next/router';
import { AuthGuard } from '@/feature/auth/secure/AuthGuard';

type GetMovieResponseResponse = {
  MovieId: number;
  Title: string;
  ReleaseYear: number;
  Description: string;
  TypeName: string;
  MovieFormat: string;
};

/**
 * TODO: 動画の詳細情報を表示する処理を追加する
 */

const MovieDetailPage = () => {
  const router = useRouter();
  const { id, movieId } = router.query;
  return (
    <AuthGuard>
      <div>
        <h1>Movie Detail Page</h1>
        <p>ID: {id}</p>
        <p>Movie ID: {movieId}</p>
      </div>
    </AuthGuard>
  );
};

export default MovieDetailPage;
