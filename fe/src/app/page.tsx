'use client'
import {Table, TableBody, TableCell, TableHead, TableHeader, TableRow} from '@/components/ui/table'
import {Button} from '@/components/ui/button'
import Modal from '@/components/custom-ui/modal'
import {useState} from 'react'
import {useRouter} from 'next/navigation'

export default function Home() {
  const router = useRouter()

  const [selectDelete, setSelectDelete] = useState<undefined|string>()

  return (
    <main className="max-w-7xl bg-white w-screen h-screen flex items-center justify-center">
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
            <TableRow>
              <TableCell>1</TableCell>
              <TableCell>The Pragmatic Programmer</TableCell>
              <TableCell>1999</TableCell>
              <TableCell>
                <Button size="icon" variant="destructive">
                  <span className="material-symbols-outlined">delete</span>
                </Button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>2</TableCell>
              <TableCell>The Clean Coder</TableCell>
              <TableCell>2011</TableCell>
              <TableCell>
                <Button size="icon" variant="destructive">
                  <span className="material-symbols-outlined">delete</span>
                </Button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>3</TableCell>
              <TableCell>Alex Ferguson: My Autobiography</TableCell>
              <TableCell>2013</TableCell>
              <TableCell>
                <Button size="icon" variant="destructive">
                  <span className="material-symbols-outlined">delete</span>
                </Button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>4</TableCell>
              <TableCell>Alderamin on the Sky</TableCell>
              <TableCell>2012</TableCell>
              <TableCell>
                <Button size="icon" variant="destructive">
                  <span className="material-symbols-outlined">delete</span>
                </Button>
              </TableCell>
            </TableRow>
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
              <Button variant="destructive">Yes, delete it</Button>
            </section>
          </section>
        </Modal>
      }
    </main>
  )
}
