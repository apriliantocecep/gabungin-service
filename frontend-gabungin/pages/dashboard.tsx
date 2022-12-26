import { GetServerSideProps, NextPage } from "next";
import Head from "next/head";
import Link from "next/link";
import { useEffect } from "react";
import Cookies from 'js-cookie';
import { useRouter } from 'next/router'
import { getAllFamily } from "../lib/api/family";
import { ApiResponse } from "../types/api";
import { Family } from "../types/family";
import FamilyStructure from "../components/Family";
import { me } from "../lib/api/account";
import { Account } from "../types/account";
import { getUserName } from "../lib/utils/account";

interface Props {
    families: Family[];
    account: Account;
}

const Dashboard: NextPage<Props> = ({ families, account }) => {

    const router = useRouter()

    const logout = () => {
        Cookies.remove("token")
        router.push("/")
    }

    useEffect(() => {
        if (!Cookies.get('token')) {
            router.push('/')
        }
    })

    return (
        <>
            <Head>
                <title>Dashboard</title>
            </Head>
            <div className="container mx-auto">
                <div className="flex items-center gap-5">
                    <a className="p-2 bg-zinc-600 text-white">Kelola Family</a>
                    <Link href={`/account`} className="p-2 bg-zinc-600 text-white">
                        Kelola Akun
                    </Link>
                    <button onClick={logout} className="p-2 bg-red-600 text-white">Keluar</button>
                </div>
                {account && (
                    <div className="flex justify-center gap-10">
                        <div className="text-center">
                            <div className={`inline-flex flex-col p-3 text-center ${account.gender == 'Male' ? 'bg-purple-800' : 'bg-pink-800'}`}>
                                <p>{getUserName(account)}</p>
                                <p>{account.gender == 'Male' ? 'Kakek' : 'Nenek'}</p>
                            </div>
                        </div>
                    </div>
                )}
                <div className="py-10 flex justify-center gap-10">
                    {families.map((family, i) => {
                        return (
                            <FamilyStructure key={i} family={family} />
                        )
                    })}
                </div>
            </div>
        </>
    )
}

export default Dashboard

export const getServerSideProps: GetServerSideProps = async (context) => {

    var families: Family[] = []

    try {
        const request = await getAllFamily(`${context.req.cookies['token']}`)
        const response: ApiResponse<Family[]> = await request.json()
        if (response.data != null) {
            families = response.data
        }
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
            families: families,
            account: account,
        }
    }
}