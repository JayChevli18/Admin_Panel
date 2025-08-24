'use client';
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Users, LayoutDashboard } from "lucide-react";

export const Sidebar = () => {
    const pathname = usePathname();
    const links = (href: string, label: string, Icon: any) => {
        const active = pathname === href;
        return (
            <Link
                href={href}
                className={`flex items-center gap-2 px-4 py-2 rounded-md hover:bg-gray-200 ${active ? 'bg-gray-300 font-semibold' : 'text-gray-700'
                    }`}
            >
                <Icon className="w-4 h-4" />
                <span>{label}</span>
            </Link>
        )
    }
    return (
        <div className="w-64 border-r bg-white/80 backdrop-blur supports-[backdrop-filter]:bg-white/60 p-4 space-y-2">
            <div className="px-2 py-3">
                <h2 className="text-lg font-semibold tracking-tight">Admin Panel</h2>
            </div>
            <nav className="space-y-2">
                {links('/', 'Dashboard', LayoutDashboard)}
                {links('/users', 'Users', Users)}
            </nav>
        </div>
    );
};

