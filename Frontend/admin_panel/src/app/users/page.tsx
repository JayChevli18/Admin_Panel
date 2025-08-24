import { UserTable } from "@/components/user-table";

export default function UsersPage () {
    return (
        <div className="space-y-4">
            <h1 className="text-2xl font-semibold tracking-tight">Users</h1>
            <UserTable />
        </div>
    );
}