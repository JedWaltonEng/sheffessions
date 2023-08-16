export default function Home() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center justify-center p-24">
        <div className="z-10 max-w-5xl w-full flex flex-col items-center justify-center font-mono text-sm lg:flex">
          <img
            src="/logo.png"
            alt="Company Logo"
            className="logo"
          />
          <p className="mb-4 text-xl font-bold">Sheffessions_</p>
          <textarea
            color="black"
            rows={5}
            placeholder="Type your anonymous confession here..."
            className="w-11/12 h-48 p-6 border-2 rounded-md text-xl mb-4 md:w-3/4 lg:w-2/3"
          ></textarea>
          <button className="p-2 bg-blue-500 text-white rounded-md">
            Submit
          </button>
        </div>
      </main>
    </>
  );
}

