'use client'
import {Form, FormControl, FormField, FormItem, FormMessage} from '@/components/ui/form'
import {z} from 'zod'
import {useForm} from 'react-hook-form'
import {zodResolver} from '@hookform/resolvers/zod'
import {Label} from '@/components/ui/label'
import {Input} from '@/components/ui/input'
import {Button} from '@/components/ui/button'
import {toastError, toastSuccess} from '@/components/custom-ui/notification'
import {Card, CardContent, CardHeader, CardTitle} from '@/components/ui/card'
import {postData} from '@/lib/fetch'
import {useRouter} from 'next/navigation'


const formSchema = z.object({
  title: z.string().min(1),
  yop: z.string().regex(/^\d{1,4}$/, {message: "Must only contain number between 0 to 9999"}),
  author: z.string().min(1),
  page: z.string(),
  isbn: z.string().regex(/^(?:|\d{10}|\d{13})$/)
})
export default function NewBook() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      title: "",
      yop: "",
      author: "",
      page: "",
      isbn: "",
    },
  })

  const router = useRouter()

  const submitForm = async (param: z.infer<typeof formSchema>) => {
    const payload = {
      title: form.getValues('title'),
      yop: parseInt(form.getValues('yop')),
      author: form.getValues('author'),
      page: form.getValues('page').length < 1 ? undefined : parseInt(form.getValues('page')),
      isbn: form.getValues('isbn').length < 1 ? undefined : form.getValues('isbn'),
    }
    try {
      const res = await postData(`${process.env.apiUrl}/books`, payload)
      toastSuccess(res)
      router.replace('/')
    } catch (err: any) {
      toastError(err)
    }
  }

  return (
    <section className="my-8 min-w-[480px]">
      <Card>
        <CardHeader>
          <CardTitle>Add new book to the library</CardTitle>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(submitForm)}>
              <FormField
                control={form.control}
                name="title"
                render={({field}) => (
                  <FormItem className="mb-2">
                    <Label htmlFor="input-title">Title of the Book <span className="text-red-500">*</span></Label>
                    <FormControl>
                      <Input id="input-title" placeholder="Enter the book title" type="text" {...field} />
                    </FormControl>
                    <div className="h-4">
                      <FormMessage className="font-normal" />
                    </div>
                  </FormItem>
                )}/>
              <FormField
                control={form.control}
                name="yop"
                render={({field}) => (
                  <FormItem className="mb-2">
                    <Label htmlFor="input-yop">Year of Publication <span className="text-red-500">*</span></Label>
                    <FormControl>
                      <Input id="input-yop" placeholder="Enter the book's year of publicatiop" type="text" {...field} />
                    </FormControl>
                    <div className="h-4">
                      <FormMessage className="font-normal" />
                    </div>
                  </FormItem>
                )}/>
              <FormField
                control={form.control}
                name="author"
                render={({field}) => (
                  <FormItem className="mb-2">
                    <Label htmlFor="input-author">Author <span className="text-red-500">*</span></Label>
                    <FormControl>
                      <Input id="input-author" placeholder="Enter the author of the book" type="text" {...field} />
                    </FormControl>
                    <div className="h-4">
                      <FormMessage/>
                    </div>
                  </FormItem>
                )}/>
              <FormField
                control={form.control}
                name="page"
                render={({field}) => (
                  <FormItem className="mb-2">
                    <Label htmlFor="input-page">Total Pages</Label>
                    <FormControl>
                      <Input id="input-page" placeholder="Enter the total pages of the book" type="text" {...field} />
                    </FormControl>
                    <div className="h-4">
                      <FormMessage className="font-normal" />
                    </div>
                  </FormItem>
                )}/>
              <FormField
                control={form.control}
                name="isbn"
                render={({field}) => (
                  <FormItem>
                    <Label htmlFor="input-isbn">ISBN</Label>
                    <FormControl>
                      <Input id="input-isbn" placeholder="Enter 10 or 13 ISBN number" type="password" {...field} />
                    </FormControl>
                    <div className="h-4">
                      <FormMessage className=" font-normal" />
                    </div>
                  </FormItem>
                )}/>
              <Button type="submit" className="mt-8">Daftar</Button>
            </form>
          </Form>
        </CardContent>
      </Card>
    </section>
  )
}