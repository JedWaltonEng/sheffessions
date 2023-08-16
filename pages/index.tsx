
export default function Home() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center justify-center p-24">
        <div className="z-10 max-w-5xl w-full flex flex-col items-center justify-center font-mono text-sm lg:flex">
          <p className="mb-4 text-xl font-bold">Sheffessions_</p>
          <textarea
            rows={5}
            placeholder="Type your anonymous confession here..."
            className="w-full md:w-1/2 lg:w-1/1.5 h-24 p-4 border-2 rounded-md text-lg mb-4"
          ></textarea>

          <button className="p-2 bg-blue-500 text-white rounded-md">
            Submit
          </button>
        </div>
      </main>
    </>
  );
}

