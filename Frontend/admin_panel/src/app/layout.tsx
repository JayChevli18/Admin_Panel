import type { Metadata } from "next";
import './globals.css';
import { Sidebar } from "@/components/sidebar";
import { ToastContainer } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';

export const metadata: Metadata = {
  title: 'Admin Panel',
  description: 'A simple admin panel built with Next.js, React Hook Form, and Tailwind CSS.',
};

export default function RootLayout({children}: {children: React.ReactNode}) {
  return (
    <html lang="en">
      <body>
        <div className="flex min-h-screen bg-gray-100">
          <Sidebar />
          <main className="flex-1 p-6">
            {children}
          </main>
        </div>
        <ToastContainer position="top-right" autoClose={3000} hideProgressBar={false} newestOnTop={false} closeOnClick rtl={false} pauseOnFocusLoss draggable pauseOnHover theme="colored"/>
      </body>
    </html>
  );
}