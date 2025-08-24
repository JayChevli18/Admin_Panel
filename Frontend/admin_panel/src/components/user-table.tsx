'use client';
import * as React from "react";
import { useState, useMemo } from "react";
import { ColumnDef, flexRender, getCoreRowModel, getPaginationRowModel, useReactTable } from "@tanstack/react-table";
import { User } from "@/types/user";
import { seedUsers } from "@/lib/data";
import { Button } from "./ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { UserForm } from './user-form';
import { ConfirmDialog } from './confirm-dialog';
import { toast } from 'react-toastify';
import Link from 'next/link';

export const UserTable = () => {
    const [data, setData] = useState<User[]>(seedUsers);
    const [editing, setEditing] = useState<User | null>(null);
    const [deleting, setDeleting] = useState<User | null>(null);

    const columns = useMemo<ColumnDef<User>[]>(() => [
        { header: 'ID', accessorKey: 'id' },
        { header: 'Name', accessorFn: (row) => `${row.firstName} ${row.lastName}` },
        { header: 'Email', accessorKey: 'email' },
        { header: 'Role', accessorKey: 'role' },
        { header: 'Status', accessorFn: (row) => (row.isActive ? 'Active' : 'Inactive') },
        {
            header: 'Actions',
            id: 'actions',
            cell: ({ row }) => {
                const user = row.original;
                return (
                    <div className="flex gap-2">
                        <Button size="sm" variant="outline" onClick={() => setEditing(user)}>Edit</Button>
                        <Button size="sm" variant="destructive" onClick={() => setDeleting(user)}>Delete</Button>
                    </div>
                )
            }
        }
    ], []);

    const table = useReactTable({
        data,
        columns,
        getCoreRowModel: getCoreRowModel(),
        getPaginationRowModel: getPaginationRowModel(),
        initialState: { pagination: { pageIndex: 0, pageSize: 5 } },
    });

    const handleUpdate = (values: any) => {
        if (!editing) return;
        setData((prev) => prev.map((u) => u.id === editing.id ? { ...editing, ...values } : u));
        setEditing(null);
        toast.success('User updated successfully');
    }

    const handleDelete = () => {
        if (!deleting) return;
        setData((prev) => prev.filter((u) => u.id !== deleting.id));
        setDeleting(null);
        toast.success('User deleted successfully');
    }

    return (
        <Card className="shadow-sm">
            <CardHeader className="flex flex-row items-center justify-between">
                <CardTitle className="text-xl">Users</CardTitle>
                <Link href="/users/new"><Button>Add User</Button></Link>
            </CardHeader>
            <CardContent>
                <div className="overflow-x-auto rounded-lg border">
                    <table className="w-full bg-white text-sm">
                        <thead className="bg-slate-50">
                            {table.getHeaderGroups().map((hg) => (
                                <tr key={hg.id} className="text-left">
                                    {hg.headers.map((h) => (
                                        <th key={h.id} className="px-4 py-3 font-medium text-slate-700">
                                            {h.isPlaceholder ? null : flexRender(h.column.columnDef.header, h.getContext())}
                                        </th>
                                    ))}
                                </tr>
                            ))}
                        </thead>
                        <tbody>
                            {table.getRowModel().rows.map((row) => (
                                <tr key={row.id} className="border-t text-left">
                                    {row.getVisibleCells().map((cell) => (
                                        <td key={cell.id} className="px-4 py-3">
                                            {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                        </td>
                                    ))}
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>


                {/* Pagination */}
                <div className="mt-4 flex items-center justify-between gap-2">
                    <div className="text-xs text-slate-500">
                        Page {table.getState().pagination.pageIndex + 1} of {table.getPageCount()}
                    </div>
                    <div className="flex items-center gap-2">
                        <Button size="sm" variant="outline" onClick={() => table.previousPage()} disabled={!table.getCanPreviousPage()}>
                            Previous
                        </Button>
                        <Button size="sm" variant="outline" onClick={() => table.nextPage()} disabled={!table.getCanNextPage()}>
                            Next
                        </Button>
                        <select
                            className="ml-2 rounded-md border border-slate-200 bg-white px-2 py-1 text-sm"
                            value={table.getState().pagination.pageSize}
                            onChange={(e) => table.setPageSize(Number(e.target.value))}
                        >
                            {[5, 10, 20].map((s) => (
                                <option key={s} value={s}>{s} / page</option>
                            ))}
                        </select>
                    </div>
                </div>
            </CardContent>
            
            {/* Edit Dialog */}
            <Dialog open={!!editing} onOpenChange={(v) => !v && setEditing(null)}>
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Edit user</DialogTitle>
                    </DialogHeader>
                    <UserForm
                        defaultValues={editing || undefined}
                        onSubmit={handleUpdate}
                        submitLabel="Update"
                    />
                </DialogContent>
            </Dialog>

            {/* Delete Confirm */}
            <ConfirmDialog
                isOpen={!!deleting}
                title="Delete user"
                description={`Are you sure you want to delete ${deleting?.firstName} ${deleting?.lastName}?`}
                onCancel={() => setDeleting(null)}
                onConfirm={handleDelete}
            />
        </Card>

    )
}