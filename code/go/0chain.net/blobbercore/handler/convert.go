package handler

import (
	"context"
	"time"

	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/allocation"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/blobbergrpc"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/datastore"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/reference"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/stats"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/writemarker"
	"github.com/0chain/blobber/code/go/0chain.net/core/common"
)

func AllocationToGRPCAllocation(alloc *allocation.Allocation) *blobbergrpc.Allocation {
	if alloc == nil {
		return nil
	}

	terms := make([]*blobbergrpc.Term, 0, len(alloc.Terms))
	for _, t := range alloc.Terms {
		terms = append(terms, &blobbergrpc.Term{
			ID:           t.ID,
			BlobberID:    t.BlobberID,
			AllocationID: t.AllocationID,
			ReadPrice:    t.ReadPrice,
			WritePrice:   t.WritePrice,
		})
	}
	return &blobbergrpc.Allocation{
		ID:               alloc.ID,
		Tx:               alloc.Tx,
		TotalSize:        alloc.TotalSize,
		UsedSize:         alloc.UsedSize,
		OwnerID:          alloc.OwnerID,
		OwnerPublicKey:   alloc.OwnerPublicKey,
		Expiration:       int64(alloc.Expiration),
		AllocationRoot:   alloc.AllocationRoot,
		BlobberSize:      alloc.BlobberSize,
		BlobberSizeUsed:  alloc.BlobberSizeUsed,
		LatestRedeemedWM: alloc.LatestRedeemedWM,
		IsRedeemRequired: alloc.IsRedeemRequired,
		TimeUnit:         int64(alloc.TimeUnit),
		CleanedUp:        alloc.CleanedUp,
		Finalized:        alloc.Finalized,
		Terms:            terms,
		PayerID:          alloc.PayerID,
	}
}

func GRPCAllocationToAllocation(alloc *blobbergrpc.Allocation) *allocation.Allocation {
	if alloc == nil {
		return nil
	}

	terms := make([]*allocation.Terms, 0, len(alloc.Terms))
	for _, t := range alloc.Terms {
		terms = append(terms, &allocation.Terms{
			ID:           t.ID,
			BlobberID:    t.BlobberID,
			AllocationID: t.AllocationID,
			ReadPrice:    t.ReadPrice,
			WritePrice:   t.WritePrice,
		})
	}
	return &allocation.Allocation{
		ID:               alloc.ID,
		Tx:               alloc.Tx,
		TotalSize:        alloc.TotalSize,
		UsedSize:         alloc.UsedSize,
		OwnerID:          alloc.OwnerID,
		OwnerPublicKey:   alloc.OwnerPublicKey,
		Expiration:       common.Timestamp(alloc.Expiration),
		AllocationRoot:   alloc.AllocationRoot,
		BlobberSize:      alloc.BlobberSize,
		BlobberSizeUsed:  alloc.BlobberSizeUsed,
		LatestRedeemedWM: alloc.LatestRedeemedWM,
		IsRedeemRequired: alloc.IsRedeemRequired,
		TimeUnit:         time.Duration(alloc.TimeUnit),
		CleanedUp:        alloc.CleanedUp,
		Finalized:        alloc.Finalized,
		Terms:            terms,
		PayerID:          alloc.PayerID,
	}
}

func FileStatsToFileStatsGRPC(fileStats *stats.FileStats) *blobbergrpc.FileStats {
	if fileStats == nil {
		return nil
	}

	return &blobbergrpc.FileStats{
		ID:                       fileStats.ID,
		RefID:                    fileStats.RefID,
		NumUpdates:               fileStats.NumUpdates,
		NumBlockDownloads:        fileStats.NumBlockDownloads,
		SuccessChallenges:        fileStats.SuccessChallenges,
		FailedChallenges:         fileStats.FailedChallenges,
		LastChallengeResponseTxn: fileStats.LastChallengeResponseTxn,
		WriteMarkerRedeemTxn:     fileStats.WriteMarkerRedeemTxn,
		CreatedAt:                fileStats.CreatedAt.UnixNano(),
		UpdatedAt:                fileStats.UpdatedAt.UnixNano(),
	}
}

func WriteMarkerToWriteMarkerGRPC(wm *writemarker.WriteMarker) *blobbergrpc.WriteMarker {
	if wm == nil {
		return nil
	}

	return &blobbergrpc.WriteMarker{
		AllocationRoot:         wm.AllocationRoot,
		PreviousAllocationRoot: wm.PreviousAllocationRoot,
		AllocationID:           wm.AllocationID,
		Size:                   wm.Size,
		BlobberID:              wm.BlobberID,
		Timestamp:              int64(wm.Timestamp),
		ClientID:               wm.ClientID,
		Signature:              wm.Signature,
	}
}

func WriteMarkerGRPCToWriteMarker(wm *blobbergrpc.WriteMarker) *writemarker.WriteMarker {
	if wm == nil {
		return nil
	}

	return &writemarker.WriteMarker{
		AllocationRoot:         wm.AllocationRoot,
		PreviousAllocationRoot: wm.PreviousAllocationRoot,
		AllocationID:           wm.AllocationID,
		Size:                   wm.Size,
		BlobberID:              wm.BlobberID,
		Timestamp:              common.Timestamp(wm.Timestamp),
		ClientID:               wm.ClientID,
		Signature:              wm.Signature,
	}
}

func FileStatsGRPCToFileStats(fileStats *blobbergrpc.FileStats) *stats.FileStats {
	if fileStats == nil {
		return nil
	}

	return &stats.FileStats{
		ID:                       fileStats.ID,
		RefID:                    fileStats.RefID,
		NumUpdates:               fileStats.NumUpdates,
		NumBlockDownloads:        fileStats.NumBlockDownloads,
		SuccessChallenges:        fileStats.SuccessChallenges,
		FailedChallenges:         fileStats.FailedChallenges,
		LastChallengeResponseTxn: fileStats.LastChallengeResponseTxn,
		WriteMarkerRedeemTxn:     fileStats.WriteMarkerRedeemTxn,
		ModelWithTS: datastore.ModelWithTS{
			CreatedAt: time.Unix(0, fileStats.CreatedAt),
			UpdatedAt: time.Unix(0, fileStats.UpdatedAt),
		},
	}
}

func CollaboratorToGRPCCollaborator(c *reference.Collaborator) *blobbergrpc.Collaborator {
	if c == nil {
		return nil
	}

	return &blobbergrpc.Collaborator{
		RefId:     c.RefID,
		ClientId:  c.ClientID,
		CreatedAt: c.CreatedAt.UnixNano(),
	}
}

func GRPCCollaboratorToCollaborator(c *blobbergrpc.Collaborator) *reference.Collaborator {
	if c == nil {
		return nil
	}

	return &reference.Collaborator{
		RefID:     c.RefId,
		ClientID:  c.ClientId,
		CreatedAt: time.Unix(0, c.CreatedAt),
	}
}

func ReferencePathToReferencePathGRPC(recursionCount *int, refPath *ReferencePath) *blobbergrpc.ReferencePath {
	if refPath == nil {
		return nil
	}
	// Accounting for bad reference paths where child path points to parent path and causes this algorithm to never end
	*recursionCount += 1
	defer func() {
		*recursionCount -= 1
	}()

	if *recursionCount > 150 {
		return &blobbergrpc.ReferencePath{
			MetaData: reference.FileRefToFileRefGRPC(reference.ListingDataToRef(refPath.Meta)),
			List:     nil,
		}
	}

	var list []*blobbergrpc.ReferencePath
	for i := range refPath.List {
		list = append(list, ReferencePathToReferencePathGRPC(recursionCount, refPath.List[i]))
	}

	return &blobbergrpc.ReferencePath{
		MetaData: reference.FileRefToFileRefGRPC(reference.ListingDataToRef(refPath.Meta)),
		List:     list,
	}
}

func ReferencePathGRPCToReferencePath(recursionCount *int, refPath *blobbergrpc.ReferencePath) *ReferencePath {
	if refPath == nil {
		return nil
	}
	// Accounting for bad reference paths where child path points to parent path and causes this algorithm to never end
	*recursionCount += 1
	defer func() {
		*recursionCount -= 1
	}()

	if *recursionCount > 150 {
		return &ReferencePath{
			Meta: reference.FileRefGRPCToFileRef(refPath.MetaData).GetListingData(context.Background()),
			List: nil,
		}
	}

	var list []*ReferencePath
	for i := range refPath.List {
		list = append(list, ReferencePathGRPCToReferencePath(recursionCount, refPath.List[i]))
	}

	return &ReferencePath{
		Meta: reference.FileRefGRPCToFileRef(refPath.MetaData).GetListingData(context.Background()),
		List: list,
	}
}
