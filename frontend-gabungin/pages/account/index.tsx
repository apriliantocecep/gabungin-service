import { GetServerSideProps, NextPage } from "next";
import { useRouter } from "next/router";
import { ParsedUrlQuery } from "querystring";
import { useState } from "react";
import { browseAccount, me } from "../../lib/api/account";
import { Account } from "../../types/account";
import { ApiResponse } from "../../types/api";
import ReactPaginate from 'react-paginate';

interface Props {
    account: Account;
    accounts: Account[];
    page: number;
    pageSize: number;
}

interface Query extends ParsedUrlQuery {
    page: string;
}

const AccountPage: NextPage<Props> = ({ account, accounts, page, pageSize }) => {

    const router = useRouter()


    return (
        <div className="container mx-auto">
            <h1 className="text-3xl">Account</h1>
            <table className="border">
                <thead>
                    <tr className="border-b">
                        <th className="border-r p-2">Id</th>
                        <th className="border-r p-2">First Name</th>
                        <th className="border-r p-2">Last Name</th>
                        <th className="border-r p-2">Email</th>
                        <th className="border-r p-2">Username</th>
                        <th className="border-r p-2">Gender</th>
                        <th className="border-r p-2">Status</th>
                        <th className="border-r p-2">Action</th>
                    </tr>
                </thead>
                <tbody>
                    {accounts.map((acc, i) => {
                        return (
                            <tr key={i}>
                                <td className="border-r border-b p-2">{acc.id}</td>
                                <td className="border-r border-b p-2">{acc.first_name}</td>
                                <td className="border-r border-b p-2">{acc.last_name}</td>
                                <td className="border-r border-b p-2">{acc.email}</td>
                                <td className="border-r border-b p-2">{acc.username}</td>
                                <td className="border-r border-b p-2">{acc.gender}</td>
                                <td className="border-r border-b p-2">{acc.status == 1 ? 'Active' : 'Disabled'}</td>
                                <td className="border-r border-b p-2">
                                    <div className="flex flex-col gap-5">
                                        <a href="" className="text-sky-600">Edit</a>
                                        <a href="" className="text-red-600">Delete</a>
                                    </div>
                                </td>
                            </tr>
                        )
                    })}
                </tbody>
            </table>

            <div className="my-3 text-center">
                
            </div>
        </div>
    )
}

export default AccountPage

export const getServerSideProps: GetServerSideProps = async (context) => {

    // get all
    var accounts: Account[] = []
    var page: number = 1
    var pageSize: number = 1
    try {
        const request = await browseAccount(`${context.req.cookies['token']}`)
        const response: ApiResponse<Account[]> = await request.json()
        if (response.data) {
            accounts = response.data
        }
        page = response.page
        pageSize = response.pageSize
    } catch (error) {

    }

    // my account
    let account: Account | null = null
    try {
        const request = await me(`${context.req.cookies['token']}`)
        const response: ApiResponse<Account> = await request.json()
        account = response.data
    } catch (error) {
    }

    return {
        props: {
            account: account,
            accounts: accounts,
            page,
            pageSize,
        }
    }
}