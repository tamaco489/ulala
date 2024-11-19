import { AuthGuard } from '@/feature/auth/secure/AuthGuard';
import { API_HOST, API_REQUEST_OPTIONS } from '@/config/env';
import axios from 'axios';
import { useEffect, useState } from 'react';
import Link from 'next/link';
import humps from 'humps';

type GetMovieListByTypeIDResponse = {
  movieId: number;
  title: string;
  releaseYear: number;
  description: string;
  typeName: string;
  movieFormat: string;
};

const MovieListPage = () => {
  const [movieList, setMovieList] = useState<
    Array<GetMovieListByTypeIDResponse>
  >([]);

  const [typeId, setTypeId] = useState<string>('');

  function extractLastNumber(path: string) {
    const segment = path.split('/').filter((s) => s !== ''); // ["movies", "types", "1"]
    const typeId = segment[segment.length - 1];
    setTypeId(typeId);
    return typeId;
  }

  async function getMovieListByTypeID(typeId: string) {
    axios.interceptors.response.use((response: any) => {
      response.data = humps.camelizeKeys(response.data);
      return response;
    });
    try {
      const res = await axios.get<Array<GetMovieListByTypeIDResponse>>(
        `${API_HOST}/movies/type?id=${typeId}`,
        API_REQUEST_OPTIONS
      );
      if (res.data.length === 0) {
        throw new Error('No movie data');
      }
      console.log('[検証中] res.data:', res.data[0]);
      setMovieList(res.data);
    } catch (err) {
      console.error('MovieList request failed:', err);
      throw err;
    }
  }

  // useEffect外でwindowオブジェクトにアクセスすると2回目のレンダリングで "ReferenceError: window is not defined" でエラーになる
  // SSRしているためuseEffect内でのみwindowオブジェクトにアクセスさせる
  // ※参考：https://dev-k.hatenablog.com/entry/how-to-access-the-window-object-in-nextjs-dev-k
  useEffect(() => {
    const currentPath = window.location.pathname;
    getMovieListByTypeID(extractLastNumber(currentPath));
  }, []);

  return (
    <AuthGuard>
      <div>
        <h1 className="underline">Movie type {} List</h1>
        <div className="text-xs my-4">
          <p>■ここでは、type毎に関連つけられた動画一覧ページを表示します。</p>
          <p>・[x] type毎の動画情報を取得するためのAPIリクエストを送信</p>
          <p>
            ・[x]
            画面上には動画のサムネイル画像とタイトル、リリース年、動画情報を表示します。
          </p>
          <p className="text-red-500">
            ・[x]
            バックエンドからのリクエストで受け取る値のkeyをスネークケースではなく、キャメルケースに変更する
          </p>
          <p className="text-blue-500">
            ・[] 画像をクリックするとその動画の詳細ページにアクセスします。
          </p>
        </div>
        <div>
          {movieList.map((movie, idx) => (
            <Link key={idx} href={`/movies/types/${typeId}/${movie.movieId}`}>
              <div
                key={idx}
                className="text-xs my-2 p-2 bg-indigo-100 border rounded-md"
              >
                <h2 className="p-1 border-solid border-b-2 border-gray-400">
                  movie_id: {movie.movieId}
                  <span> | {movie.title}</span>
                </h2>
                <p className="py-2 px-1">リリース年: {movie.releaseYear}</p>
                <p className="px-1">説明: {movie.description}</p>
              </div>
            </Link>
          ))}
        </div>
      </div>
    </AuthGuard>
  );
};

export default MovieListPage;
