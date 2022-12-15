import React, { useCallback, useState } from "react";
import ReactDOM from "react-dom";

const App = (props: { message: string }) => {
  const [count, setCount] = useState(0);
  const increment = useCallback(() => {
    setCount((count) => count + 1);
  }, [count]);
  return (
    // From: https://www.hyperui.dev/components/marketing/banners#component-1
    <>
      <section className="bg-gray-50">
        <div className="mx-auto max-w-screen-xl px-4 py-32 lg:flex lg:h-screen lg:items-center">
          <div className="mx-auto max-w-xl text-center">
            <h1 className="text-3xl font-extrabold sm:text-5xl">
              Hello World!
              <strong className="font-extrabold text-red-700 sm:block">
                Count: {count}
              </strong>
            </h1>

            <p className="mt-4 sm:text-xl sm:leading-relaxed">
              A Simple Counter App built on ESBuild + React + Typescript +
              TailwindCSS.
            </p>

            <div className="mt-8 flex flex-wrap justify-center gap-4">
              <button
                className="block w-full rounded bg-red-600 px-12 py-3 text-sm font-medium text-white shadow hover:bg-red-700 focus:outline-none focus:ring active:bg-red-500 sm:w-auto"
                onClick={increment}
              >
                Increment
              </button>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

ReactDOM.render(
  <App message="Hello World! A Simple Counter App built on ESBuild + React + Typescript" />,
  document.getElementById("root")
);
