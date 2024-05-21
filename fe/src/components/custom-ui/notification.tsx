import {toast} from '@/components/ui/use-toast'
import {toTitleCase} from '@/lib/helper'


export function toastSuccess(toastData: any, customMsg?: string, customDesc?: string) {
  toast({
    title: customMsg || toTitleCase(toastData.message.replaceAll('_', ' ')),
    description: customDesc,
    variant: "success",
  })
}

export function toastError(toastData: any, customMsg?: string, customDesc?: string) {
  toast({
    title: toastData ? toTitleCase(toastData.message.replaceAll('_', ' ')) : customMsg,
    description: toastData ? toastData.data : customDesc,
    variant: "destructive",
  })
}