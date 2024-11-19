import { AuthGuard } from '@/feature/auth/secure/AuthGuard';
import { API_HOST, API_REQUEST_OPTIONS } from '@/config/env';
import axios from 'axios';
import Image from 'next/image';
import { useEffect, useState } from 'react';
import Link from 'next/link';
import { NOT_IMAGE_TYPE_ID } from '@/config/constants';
import humps from 'humps';

type GetMovieCategoriesResponse = {
  typeId: number;
  typeName: string;
  title: string;
  description: string;
};

// movieのカテゴリ一覧を表示するページ
const MovieCategoriesPage = () => {
  const [movieCategories, setMovieCategories] = useState<
    Array<GetMovieCategoriesResponse>
  >([]);

  const getMovieCategories = async () => {
    axios.interceptors.response.use((response: any) => {
      response.data = humps.camelizeKeys(response.data);
      return response;
    });
    try {
      const res = await axios.get<Array<GetMovieCategoriesResponse>>(
        `${API_HOST}/movies/categories`,
        API_REQUEST_OPTIONS
      );
      setMovieCategories(res.data);
    } catch (err) {
      console.error('MovieCategories request failed:', err);
      throw err;
    }
  };

  useEffect(() => {
    getMovieCategories();
  }, []); // 空の依存配列を指定することで、初回のマウント時にのみ実行される

  return (
    <AuthGuard>
      <div>
        <h2 className="text-xl underline">Movies List Page</h2>
        <div className="">
          {movieCategories.map((image, idx) =>
            image.typeId === NOT_IMAGE_TYPE_ID ? null : (
              <Link key={idx} href={`/movies/types/${image.typeId}`}>
                <div className="my-4 border border-gray-500 rounded-md">
                  <p className="text-sm font-semibold p-2 border border-gray-500 rounded-t-md">
                    {image.typeId}: {image.title} ({image.typeName})
                  </p>
                  <div className="">
                    <Image
                      src={`/proto/img/thumbnail/type${image.typeId}/${image.typeId}0000000.png`}
                      alt={image.description}
                      width={100}
                      height={100}
                      className="w-full"
                    />
                    <p className="text-xs p-2 border border-gray-500 rounded-b-md">
                      {image.description}
                    </p>
                  </div>
                </div>
              </Link>
            )
          )}
        </div>
      </div>
    </AuthGuard>
  );
};

export default MovieCategoriesPage;
