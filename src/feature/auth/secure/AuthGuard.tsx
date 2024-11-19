import { useRouter } from 'next/router';
import { useAuthContext } from '../provider/AuthProvider';

type Props = {
  children: React.ReactNode;
};

export const AuthGuard = ({ children }: Props) => {
  const { user } = useAuthContext();
  const { push } = useRouter();

  if (typeof user === 'undefined') {
    return <div className="bg-gray-100 text-center py-4">Loading...</div>;
  }

  if (user === null) {
    push('/signin');
    return null;
  }

  return <>{children}</>;
};
