import SubmissionForm from "@/components/SubmissionForm";

export default function Home() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center justify-start p-12 md:justify-center">
        <div className="z-10 max-w-5xl w-full flex flex-col items-center justify-center font-mono text-sm">
          <img
            src="/logo.png"
            alt="Company Logo"
            className="logo mb-4"
          />
          <p className="mb-4 text-xl font-bold">Sheffessions_</p>
          <SubmissionForm />
        </div>
      </main>
    </>
  );
}

