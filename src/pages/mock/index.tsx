import axios from 'axios';
import { FormEvent, useState } from 'react';

type TodoType = {
  id: number;
  userId: number;
  title: string;
  completed?: boolean;
};

export const MockPage = () => {
  const baseURL = 'https://jsonplaceholder.typicode.com';
  const [todoObj, setTodoObj] = useState<TodoType>({
    id: 0,
    userId: 0,
    title: '',
    completed: false,
  });
  const [todoAll, setTodoAll] = useState<Array<TodoType>>([]);
  const [inputTodoID, setInputTodoID] = useState('');

  const getTodoAllHandler = () => {
    const todoAllData = axios
      .get<Array<TodoType>>(`${baseURL}/todos`)
      .then((res) => {
        setTodoAll(res.data);
      });

    return todoAllData;
  };

  const getTodoObjHandler = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const parsedTodoID = parseInt(inputTodoID, 10);
    if (isNaN(parsedTodoID)) {
      return alert('Invalid input id: ' + inputTodoID);
    }

    if (200 < parsedTodoID || parsedTodoID < 1) {
      return alert('must enter integer between 1 and 200');
    }

    /* TODO:
      - 全角入力された場合に半角に変換する制御を追加
      - そもそもフリーフォーマットで受け付けているのが良くないので、プルダウン形式等で数値以外の入力ができないようにしたら良さそう
    */
    const todoObjData = axios
      .get<TodoType>(`${baseURL}/todos/${parsedTodoID}`)
      .then((res) => {
        setTodoObj(res.data);
      })
      .catch((err: any) => {
        console.error('request failed', err);
      });

    return todoObjData;
  };

  const resetDataHandler = () => {
    setTodoAll([]);
    setTodoObj({ id: 0, userId: 0, title: '', completed: false });
  };

  return (
    <div className="my-4">
      <h2 className="mt-8 mb-4 text-center text-bold">Mock Page</h2>
      <div className="grid gap-4">
        <div>
          <button
            className="px-4 py-1 mr-2 text-white bg-red-500 rounded"
            onClick={resetDataHandler}
          >
            reset
          </button>
        </div>

        <div className="border-2 border-blue-400">
          <p className="mb-2 text-white font-bold bg-blue-400">
            Get Todo Specify UserID
          </p>
          <div className="flex">
            <form onSubmit={(e) => getTodoObjHandler(e)}>
              <label className="block px-2">input todo id</label>
              <input
                type="text"
                value={inputTodoID}
                className="mx-2 border border-gray-300"
                onChange={(e) => {
                  setInputTodoID(e.target.value);
                }}
              />
              <button
                type="submit"
                className="my-1 py-1 px-4 rounded bg-gray-500 text-white"
              >
                submit
              </button>
            </form>
          </div>
          <div className="m-2">
            {todoObj && (
              <div>
                <p>{`id: ${todoObj.id}`}</p>
                <p>{`user_id: ${todoObj.userId}`}</p>
                <p>{`title: ${todoObj.title || 'not contents'}`}</p>
                <p>
                  {`completed: ${
                    todoObj.completed !== undefined
                      ? todoObj.completed
                        ? 'true'
                        : 'false'
                      : 'not completed'
                  }`}
                </p>
              </div>
            )}
          </div>
        </div>

        <div className="border-2 border-green-400">
          <p className="mb-2 text-white font-bold bg-green-400">Get Todo All</p>
          <div>
            <button
              className="mx-2 py-1 px-4 rounded bg-gray-500 text-white"
              onClick={getTodoAllHandler}
            >
              submit
            </button>

            <p className="m-2 font-bold">{`object len: ${todoAll.length}`}</p>
            {todoAll.map((todo) => (
              <li
                key={todo.id}
                className="list-none mt-2 p-2 bg-green-100 rounded"
              >
                <p>{`id: ${todo.id}`}</p>
                <p>{`user_id: ${todo.userId}`}</p>
                <p>{`title: ${todo.title}`}</p>
                <p>
                  {`completed: ${
                    todo.completed !== undefined
                      ? todo.completed
                        ? 'true'
                        : 'false'
                      : 'not completed'
                  }`}
                </p>
              </li>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default MockPage;
