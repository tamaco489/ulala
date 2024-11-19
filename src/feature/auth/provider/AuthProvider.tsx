import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from 'react';
import type { User } from '@firebase/auth';
import { getAuth, onAuthStateChanged } from '@firebase/auth';

export type GlobalAuthState = {
  user:
    | User
    | null
    | undefined /* null: 未認証の場合、undefind: 初期状態(認証前) */;
};

const initiallState: GlobalAuthState = {
  user: undefined /* initiallStateは初期状態としてundefind型であることを定義 */,
};

/* createContext: 親→子へのコンポーネント伝達を実現する、引数として初期値を渡すことが可能（以下はinitiallStateを初期値として設定している） */
const AuthContext = createContext<GlobalAuthState>(initiallState);

/* ReactNode: 子コンポーネントが他要素、または他コンポーネントを受け入れための型情報をPropsとして定義（型情報: 文字列、数値、配列、フラグメント等）*/
type Props = { children: ReactNode };

export const AuthProvider = ({ children }: Props) => {
  const [user, setUser] = useState<GlobalAuthState>(initiallState);

  useEffect(() => {
    try {
      const auth = getAuth();
      return onAuthStateChanged(auth, (user) => {
        setUser({ user });
      });
    } catch (e: any) {
      setUser(initiallState);
      throw e;
    }
  }, []); // コンポーネントがページロード等によりDOMに追加（マウント）された時にのみ実行、再読み込み時にも通信が走る

  return <AuthContext.Provider value={user}>{children}</AuthContext.Provider>;
};

export const useAuthContext = () => useContext(AuthContext);
