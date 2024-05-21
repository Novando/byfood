import {Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue} from '@/components/ui/select'
import {
  Pagination,
  PaginationContent,
  PaginationIconOnly,
  PaginationItem,
  PaginationLink,
} from '@/components/ui/pagination'
import {Dispatch, SetStateAction, useCallback, useEffect} from 'react'
import {usePathname, useRouter, useSearchParams} from 'next/navigation'


export type DataTableFooterData = {
    total: number
    page: number
    size: number
}

type DataTableFooterGetter = {
  footerData: DataTableFooterData
}

type DataTableFooterSetter = {
  setFooterData: Dispatch<SetStateAction<DataTableFooterData>>
  setPageData: Dispatch<string>
}

export function DataTableFooter(props: DataTableFooterGetter & DataTableFooterSetter) {
  const searchParams = useSearchParams()
  const pathname = usePathname()
  const router = useRouter()

  useEffect(() => {
    const size = parseInt(searchParams.get('size') || '0')
    const page = parseInt(searchParams.get('page') || '1')
    const newSize = size < 10 ? 10 : (size > 100 ? 100 : size)
    const maxPage =  props.footerData.total / newSize
    props.setFooterData((prev) => {
      return {
        ...prev,
        size: newSize,
        page: page < 1 ? 1 : (page > maxPage ? maxPage : page),
      }
    })
  }, [])

  const appendQueryString = useCallback((name: string, value: string) => {
    if (name === 'size') props.setFooterData((prev: DataTableFooterData) => {
      return {...prev, size: parseInt(value)}
    })
    if (name === 'page') props.setFooterData((prev: DataTableFooterData) => {
      return {...prev, page: parseInt(value)}
    })
    const params = new URLSearchParams(searchParams.toString())
    params.set(name, value)
    router.replace(`${pathname}?${params.toString()}`)
    props.setPageData(params.toString())
  },[searchParams])


  return (
    <>
      {props.footerData.total > 0 &&
        <section className="flex items-center justify-between mx-10 mt-4">
          <p className="text-neutral-500 text-sm">
            showing { (props.footerData.page - 1) * props.footerData.size + 1 } till { props.footerData.total < (props.footerData.page * props.footerData.size) ? props.footerData.total : props.footerData.page * props.footerData.size }, total {props.footerData.total} data
          </p>
          <div className="flex items-center gap-4">
            <p>Data size</p>
            <Select onValueChange={(e: string) => appendQueryString('size', e)}>
              <SelectTrigger className="w-20">
                <SelectValue placeholder={props.footerData.size} />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <SelectItem value="10" >10</SelectItem>
                  <SelectItem value="50">50</SelectItem>
                  <SelectItem value="100">100</SelectItem>
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>
        </section>
      }
      {props.footerData.total > props.footerData.size &&
        <Pagination>
          <PaginationContent>
            <PaginationItem>
              <PaginationIconOnly hidden={props.footerData.page <= 1} linkClick={() => appendQueryString('page', '1')} size="default">
                <span className="material-symbols-outlined">first_page</span>
              </PaginationIconOnly>
            </PaginationItem>
            <PaginationItem>
              <PaginationIconOnly hidden={props.footerData.page <= 1} linkClick={() => appendQueryString('page', (props.footerData.page - 1).toString())} size="default">
                <span className="material-symbols-outlined">chevron_left</span>
              </PaginationIconOnly>
            </PaginationItem>
            <PaginationItem>
              <PaginationLink size="default">{props.footerData.page}</PaginationLink>
            </PaginationItem>
            <PaginationItem>
              <PaginationIconOnly hidden={props.footerData.page >= (props.footerData.total / props.footerData.size)} linkClick={() => appendQueryString('page', (props.footerData.page + 1).toString())} size="default">
                <span className="material-symbols-outlined">chevron_right</span>
              </PaginationIconOnly>
            </PaginationItem>
            <PaginationItem>
              <PaginationIconOnly hidden={props.footerData.page >= (props.footerData.total / props.footerData.size)} linkClick={() => appendQueryString('page', (props.footerData.total / props.footerData.size).toString())} size="default">
                <span className="material-symbols-outlined">last_page</span>
              </PaginationIconOnly>
            </PaginationItem>
          </PaginationContent>
        </Pagination>
      }
    </>
  )
}