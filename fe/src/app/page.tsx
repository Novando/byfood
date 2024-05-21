'use client'
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from '@/components/ui/table'
import {Button} from '@/components/ui/button'
import Modal from '@/components/custom-ui/modal'
import {useEffect, useState} from 'react'
import {useRouter} from 'next/navigation'
import {toastError, toastSuccess} from '@/components/custom-ui/notification'
import {delData, getData} from '@/lib/fetch'

type BookType = {
  id: string
  title: string
  yop: number
  author: string
}

export default function Home() {
  const [selectDelete, setSelectDelete] = useState<undefined|string>()
  const [books, setBooks] = useState<BookType[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const router = useRouter()

  useEffect(() => {
    getBooks()
  }, [])

  const getBooks = async () => {
    try {
      const res = await getData(`${process.env.apiUrl}/books`)
      setBooks(res.data || [])
    } catch (err) {
      toastError(err)
    }
  }

  const deleteBook = async () => {
    setIsLoading(true)
    try {
      await delData(`${process.env.apiUrl}/books/${selectDelete}`)
      const res = await getData(`${process.env.apiUrl}/books`)
      setBooks(res.data)
      toastSuccess({}, "Book deleted!")
      setSelectDelete(undefined)
    } catch (err) {
      toastError(err)
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <main className="max-w-7xl bg-white w-screen flex justify-center my-20">
      <section className="flex flex-col w-[1080px] bg-slate-400 p-8 rounded-3xl cursor-default">
        <section className="self-end mb-4">
          <Button onClick={() => router.push("/new-book")}><span className="material-symbols-outlined">add</span> Add Book</Button>
        </section>
        <Table className="bg-slate-200 rounded-xl">
          <TableHeader>
            <TableRow>
              <TableHead className="w-8">No</TableHead>
              <TableHead>Title</TableHead>
              <TableHead>Year of Publication</TableHead>
              <TableHead>Action</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {books.map((item, key) =>
              <TableRow key={item.id}>
                <TableCell>{key + 1}</TableCell>
                <TableCell onClick={() => router.push(`/${item.id}`)} className="text-sky-700">{item.title}</TableCell>
                <TableCell>{item.yop}</TableCell>
                <TableCell>
                  <Button onClick={() => setSelectDelete(item.id)} size="icon" variant="destructive">
                    <span className="material-symbols-outlined">delete</span>
                  </Button>
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </section>
      {selectDelete &&
        <Modal
          closeModal={() => setSelectDelete(undefined)}
          title="Confirm Delete"
        >
          <section>
            <p>Are you sure want to delete BOOKSNAME from the library?</p>
            <section className="flex items-center justify-center gap-8">
              <Button onClick={() => setSelectDelete(undefined)} variant="secondary">No, nevermind</Button>
              <Button onClick={deleteBook} disabled={isLoading} variant="destructive">Yes, delete it</Button>
            </section>
          </section>
        </Modal>
      }
    </main>
  )
}
